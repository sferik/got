package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/kurrik/oauth1a"
)

var programName string
var programVersion = "3.0.0"

func init() {
	programName = filepath.Base(os.Args[0])
}

func main() {
	commands := map[string]command{
		"authorize": makeAuthorizeCommand(),
		"ruler":     makeRulerCommand(),
		"version":   makeVersionCommand(),
	}

	flag.Usage = func() {
		fmt.Printf("%s is a command-line interface to Twitter.\n\n", programName)
		var maxCommandLen int
		for name := range commands {
			if len := len(name); len > maxCommandLen {
				maxCommandLen = len
			}
		}
		for name, subcommand := range commands {
			formatString := fmt.Sprintf("%%%ds: %%s\n", maxCommandLen)
			fmt.Printf(formatString, name, subcommand.desc)
		}
	}

	flag.Parse()

	args := flag.Args()

	if len(args) <= 0 || args[0] == "help" {
		flag.Usage()
		os.Exit(2)
	}
	cmd, ok := commands[args[0]]
	if !ok {
		printError("unknown subcommand %q\nRun '%s help' for usage.\n", args[0], programName)
		os.Exit(1)
	}
	if err := cmd.fn(args[1:]); err != nil {
		printError("error running %q: %s", args[0], err)
		os.Exit(1)
	}
}

type command struct {
	fs   *flag.FlagSet
	desc string
	fn   func([]string) error
}

func makeAuthorizeCommand() command {
	fn := func(args []string) error {
		message := fmt.Sprintf("Press [Enter] to authorize %s with Twitter.", programName)
		_, err := fmt.Println(message)
		reader := bufio.NewReader(os.Stdin)
		reader.ReadString('\n')
		service := &oauth1a.Service{
			RequestURL:   "https://api.twitter.com/oauth/request_token",
			AuthorizeURL: "https://api.twitter.com/oauth/authorize",
			AccessURL:    "https://api.twitter.com/oauth/access_token",
			ClientConfig: &oauth1a.ClientConfig{
				ConsumerKey:    "r1iPw4qOe4QJ9X7wgaXkGXRYH",                          // os.Getenv("TWITTER_CONSUMER_KEY"),
				ConsumerSecret: "fuCqog2zxrkSoAWgv7wisBeLCL2p0aGH3OeBTD4cw0VVCU8OIm", // os.Getenv("TWITTER_CONSUMER_SECRET"),
			},
			Signer: new(oauth1a.HmacSha1Signer),
		}
		httpClient := new(http.Client)
		userConfig := &oauth1a.UserConfig{}
		if err := userConfig.GetRequestToken(service, httpClient); err != nil {
			return err
		}
		if url, err := userConfig.GetAuthorizeURL(service); err != nil {
			return err
		} else {
			if _, err := exec.Command("open", url).Output(); err != nil {
				return err
			}
		}
		fmt.Print("Enter the supplied PIN: ")
		reader = bufio.NewReader(os.Stdin)
		reader.ReadString('\n')
		return err
	}
	return command{
		desc: fmt.Sprintf("authorize your Twitter account with %s", programName),
		fn:   fn,
	}
}

func makeRulerCommand() command {
	fs := flag.NewFlagSet("ruler", flag.ExitOnError)
	spacesToIndent := fs.Int("indent", 0, "the number of spaces to print before the ruler")
	fn := func(args []string) error {
		fs.Parse(args)
		_, err := fmt.Println(ruler(*spacesToIndent))
		return err
	}
	return command{
		fs:   fs,
		desc: "print a 140-character ruler",
		fn:   fn,
	}
}

func ruler(spacesToIndent int) string {
	var ruler string
	for i := 0; i < spacesToIndent; i++ {
		ruler += " "
	}
	for i := 1; i <= 140; i++ {
		if i%5 == 0 {
			ruler += "|"
		} else {
			ruler += "-"
		}
	}
	return ruler
}

func makeVersionCommand() command {
	fn := func(args []string) error {
		_, err := fmt.Println(programVersion)
		return err
	}
	return command{
		desc: fmt.Sprintf("print %s version", programName),
		fn:   fn,
	}
}

func printError(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, programName+": "+format, args...)
}

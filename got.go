package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

var programName string

func init() {
	programName = filepath.Base(os.Args[0])
}

func main() {
	commands := map[string]command{
		"ruler": makeRulerCommand(),
	}

	flag.Usage = func() {
		fmt.Printf("%s is a Twitter command-line interface.\n\n", programName)
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

	if len(args) <= 0 {
		flag.Usage()
		os.Exit(2)
	}
	cmd, ok := commands[args[0]]
	if !ok {
		printError("%q is not a valid subcommand\n", args[0])
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
		desc: "prints a 140-character ruler",
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

func printError(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, programName+": "+format, args...)
}

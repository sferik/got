package main

import (
	"flag"
	"fmt"
	"os"
)

var homedir string

func init() {
	homedir = os.Getenv("HOME")
}

func main() {
	var (
	// color   = flag.String("color", "auto", "control how color is used in output")
	// profile = flag.String("profile", filepath.Join(homedir, ".trc"), "path to RC file")
	)
	flag.Parse()

	commands := map[string]command{
		"ruler": makeRulerCommand(),
	}

	args := flag.Args()

	if len(args) <= 0 {
		flag.Usage()
		os.Exit(1)
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
	fs *flag.FlagSet
	fn func([]string) error
}

func makeRulerCommand() command {
	fs := flag.NewFlagSet("ruler", flag.ExitOnError)
	spacesToIndent := fs.Int("indent", 0, "the number of spaces to print before the ruler")
	fn := func(args []string) error {
		fs.Parse(args)
		return ruler(*spacesToIndent)
	}
	return command{
		fs: fs,
		fn: fn,
	}
}

func ruler(spacesToIndent int) error {
	var ruler string
	for i := 0; i < spacesToIndent; i++ {
		ruler += " "
	}
	for len(ruler) < 140 {
		ruler += "----|"
	}
	_, err := fmt.Println(ruler)
	return err
}

func printError(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, os.Args[0]+": "+format, args...)
}

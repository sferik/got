package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
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
		_, err := fmt.Println(ruler(*spacesToIndent))
		return err
	}
	return command{
		fs: fs,
		fn: fn,
	}
}

func ruler(spacesToIndent int) string {
	var ruler string
	for i := 0; i < spacesToIndent; i++ {
		ruler += " "
	}
	for len(ruler) < 140 {
		ruler += "----|"
	}
	return ruler
}

func printError(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, os.Args[0]+": "+format, args...)
}

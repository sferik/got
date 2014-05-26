package main

import (
	"flag"
	"fmt"
)

func ruler() {
	var spacesToIndent = flag.Int("indent", 0, "the number of spaces to print before the ruler")
	var ruler string
	for i := 0; i < *spacesToIndent; i++ {
		ruler += " "
	}
	for len(ruler) < 140 {
		ruler += "----|"
	}
	fmt.Println(ruler)
}

func main() {
	flag.Parse()
	args := flag.Args()
	switch args[0] {
	case "ruler":
		ruler()
	}
}

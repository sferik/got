package main

import (
	"fmt"
	"os"
)

var foo = os.Getenv("FOO")

func main() {
	fmt.Println(foo)
}

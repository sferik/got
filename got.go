package main

import (
  "flag"
  "fmt"
)

func ruler() {
  var ruler string
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

package main

import (
	"flag"
	"fmt"
)

func main() {
	msg := flag.String("message", "I'm a Gopher!", "A message you'd like to display")
	flag.Parse()

	fmt.Println(*msg)
}

package main

import (
	"bufio"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	conf := &Config{
		Next:     "",
		Previous: "",
	}
	repl(scanner, conf)
}

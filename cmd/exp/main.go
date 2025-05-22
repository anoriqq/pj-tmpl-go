package main

import (
	"fmt"
	"os"
)

func main() {
	if err := run(); err != nil {
		os.Exit(1)
	}
}

func run() error {
	fmt.Println("Hello World")
	return nil
}

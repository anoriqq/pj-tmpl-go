package internal

import (
	"flag"
	"os"
	"strings"
)

type cliOptions struct {
	Name string
}

func NewCLIOptions() cliOptions {
	opts := cliOptions{}

	flag.StringVar(&opts.Name, "name", "", "Name to greet")

	flag.VisitAll(func(f *flag.Flag) {
		if s := os.Getenv(strings.ToUpper(f.Name)); s != "" {
			f.Value.Set(s)
		}
	})

	flag.Parse()

	return opts
}

package internal

import (
	"flag"
	"log/slog"
	"os"
	"strings"
)

type cliOptions struct {
	env  string
	name string
}

func (o cliOptions) Env() string {
	return o.env
}

func (o cliOptions) Name() string {
	return o.name
}

func (o cliOptions) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("env", o.Env()),
		slog.String("name", o.Name()),
	)
}

func NewCLIOptions() cliOptions {
	opts := cliOptions{}

	flag.StringVar(&opts.env, "env", "", "Environment to use (dev, stg, prd)")
	flag.StringVar(&opts.name, "name", "", "Name to greet")

	flag.VisitAll(func(f *flag.Flag) {
		if s := os.Getenv(strings.ToUpper(f.Name)); s != "" {
			f.Value.Set(s)
		}
	})

	flag.Parse()

	return opts
}

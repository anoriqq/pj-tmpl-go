package cli

import (
	"flag"
	"log/slog"
	"os"
	"strings"
)

type options struct {
	env  string
	name string
}

func (o options) Env() string {
	return o.env
}

func (o options) Name() string {
	return o.name
}

func (o options) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("env", o.Env()),
		slog.String("name", o.Name()),
	)
}

// NewOptions CLIのオプションを取得する
// フラグと環境変数から値を取得する。両方に値が設定されている場合はフラグの値を採用する。
func NewOptions() options {
	opts := options{}

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

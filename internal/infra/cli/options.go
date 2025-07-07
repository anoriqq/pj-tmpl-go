package cli

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/anoriqq/pj-tmpl-go/internal/domain/env"
	"github.com/go-errors/errors"
)

type options struct {
	help bool
	env  env.Env
	port uint64
}

var _ slog.LogValuer = (options{})

func (o options) Help() bool {
	return o.help
}

func (o options) Env() env.Env {
	return o.env
}

func (o options) Port() uint64 {
	return o.port
}

func (o options) LogValue() slog.Value {
	return slog.GroupValue(
		slog.Bool("help", o.Help()),
		slog.String("env", o.Env().String()),
		slog.Uint64("port", o.Port()),
	)
}

var envNameReplacer = strings.NewReplacer("-", "_", ".", "_")

// NewOptions CLIのオプションを取得する
// フラグと環境変数から値を取得する。両方に値が設定されている場合はフラグの値を採用する。
func NewOptions(args []string) (options, error) {
	fs := flag.NewFlagSet("app", flag.ContinueOnError)

	opts := options{
		help: false,
		env:  env.LCL,
		port: 8080,
	}
	fs.BoolVar(&opts.help, "help", opts.help, "Show help message and exit")
	envUsage := fmt.Sprintf("Environment to use (%s)", strings.Join(env.EnvStrings(), ","))
	fs.Var(&opts.env, "env", envUsage)
	fs.Uint64Var(&opts.port, "port", opts.port, "Port to listen on")

	if parseErr := fs.Parse(args); parseErr != nil {
		return options{}, errors.Wrap(parseErr, 0)
	}

	fs.VisitAll(func(f *flag.Flag) {
		if f.Value.String() != f.DefValue {
			// フラグに値が指定されている場合はそれをそのまま使うのでスキップ
			return
		}

		envName := strings.ToUpper(envNameReplacer.Replace(f.Name))
		if env := os.Getenv(envName); env != "" {
			f.Value.Set(env)
		}
	})

	if opts.help {
		fmt.Fprintf(os.Stderr, "usage: cmd [flags]\n")
		fs.PrintDefaults()
	}

	return opts, nil
}

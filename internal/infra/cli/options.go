package cli

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/anoriqq/pj-tmpl-go/internal/domain/env"
	"github.com/anoriqq/pj-tmpl-go/internal/domain/port"
	"github.com/go-errors/errors"
)

type options struct {
	help bool
	env  env.Env
	port port.Port
}

var _ slog.LogValuer = (options{})

func (o options) Help() bool {
	return o.help
}

func (o options) Env() env.Env {
	return o.env
}

func (o options) Port() port.Port {
	return o.port
}

func (o options) LogValue() slog.Value {
	return slog.GroupValue(
		slog.Bool("help", o.Help()),
		slog.String("env", o.Env().String()),
		slog.String("port", o.Port().String()),
	)
}

var envNameReplacer = strings.NewReplacer("-", "_", ".", "_")

// NewOptions CLIのオプションを取得する
// フラグと環境変数から値を取得する。両方に値が設定されている場合はフラグの値を採用する。
func NewOptions(args []string) (options, error) {
	flagSet := flag.NewFlagSet("app", flag.ContinueOnError)

	opts := options{
		help: false,
		env:  env.LCL,
		port: port.New(8000),
	}
	flagSet.BoolVar(&opts.help, "help", opts.help, "Show help message and exit")

	envUsage := fmt.Sprintf("Environment to use (%s)", strings.Join(env.EnvStrings(), ","))
	flagSet.Var(&opts.env, "env", envUsage)
	flagSet.Var(&opts.port, "port", "Port to listen on")

	parseErr := flagSet.Parse(args)
	if parseErr != nil {
		return options{}, errors.Wrap(parseErr, 0)
	}

	flagSet.VisitAll(func(flg *flag.Flag) {
		if flg.Value.String() != flg.DefValue {
			// フラグに値が指定されている場合はそれをそのまま使うのでスキップ
			return
		}

		envName := strings.ToUpper(envNameReplacer.Replace(flg.Name))
		if env := os.Getenv(envName); env != "" {
			flg.Value.Set(env)
		}
	})

	if opts.help {
		fmt.Fprintf(os.Stderr, "usage: cmd [flags]\n")
		flagSet.PrintDefaults()
	}

	return opts, nil
}

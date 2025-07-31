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

// Options CLIのオプション
type Options struct {
	help *bool
	env  *env.Env
	port *port.Port
}

var _ slog.LogValuer = (Options{})

// Help ヘルプフラグを取得する
func (o Options) Help() bool {
	if o.help == nil {
		return false
	}
	return *o.help
}

// Env 環境を取得する
func (o Options) Env() env.Env {
	if o.env == nil {
		return env.LCL
	}
	return *o.env
}

// Port リスニングポートを取得する
func (o Options) Port() port.Port {
	if o.port == nil {
		return port.New(8000)
	}
	return *o.port
}

// LogValue implements [slog.LogValuer]
func (o Options) LogValue() slog.Value {
	return slog.GroupValue(
		slog.Bool("help", o.Help()),
		slog.String("env", o.Env().String()),
		slog.String("port", o.Port().String()),
	)
}

var envNameReplacer = strings.NewReplacer("-", "_", ".", "_")

// NewOptions CLIのオプションを取得する
// フラグと環境変数から値を取得する。両方に値が設定されている場合はフラグの値を採用する。
func NewOptions(args []string) (Options, error) {
	flagSet := flag.NewFlagSet("app", flag.ContinueOnError)

	opts := Options{}
	flagSet.BoolVar(opts.help, "help", opts.Help(), "Show help message and exit")

	envUsage := fmt.Sprintf("Environment to use (%s)", strings.Join(env.EnvStrings(), ","))
	flagSet.Var(opts.env, "env", envUsage)
	flagSet.Var(opts.port, "port", "Port to listen on")

	parseErr := flagSet.Parse(args)
	if parseErr != nil {
		return Options{}, errors.Wrap(parseErr, 0)
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

	if opts.Help() {
		fmt.Fprintf(os.Stderr, "usage: cmd [flags]\n")
		flagSet.PrintDefaults()
	}

	return opts, nil
}

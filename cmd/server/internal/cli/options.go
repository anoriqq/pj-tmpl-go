package cli

import (
	"flag"
	"log/slog"
	"os"
	"strings"
	"sync"
)

var (
	opts options
	mu   sync.Mutex
)

func init() {
	flag.StringVar(&opts.env, "env", "", "Environment to use (dev, stg, prd)")
	flag.Uint64Var(&opts.port, "port", 0, "Port to listen on")
	flag.StringVar(&opts.name, "name", "", "Name to greet")
}

type options struct {
	env  string
	port uint64
	name string
}

func (o options) Env() string {
	return o.env
}

func (o options) Port() uint64 {
	return o.port
}

func (o options) Name() string {
	return o.name
}

func (o options) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("env", o.Env()),
		slog.Uint64("port", o.Port()),
		slog.String("name", o.Name()),
	)
}

// NewOptions CLIのオプションを取得する
// フラグと環境変数から値を取得する。両方に値が設定されている場合はフラグの値を採用する。
func NewOptions() options {
	mu.Lock()
	flag.Parse()
	mu.Unlock()

	flag.VisitAll(func(f *flag.Flag) {
		if strings.Contains(f.Name, ".") {
			// ドットが含まれるフラグはoptionsに定義されているものではないので無視
			return
		}

		if f.Value.String() != f.DefValue {
			// フラグに値が指定されている場合はそれをそのまま使うのでスキップ
			return
		}

		if env := os.Getenv(strings.ToUpper(f.Name)); env != "" {
			f.Value.Set(env)
		}
	})

	return opts
}

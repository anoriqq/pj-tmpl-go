/*
Package cli provides a simple wrapper around the standard flag package
*/
package cli

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/anoriqq/pj-tmpl-go/internal/domain/env"
	"github.com/anoriqq/pj-tmpl-go/internal/domain/port"
	"github.com/go-errors/errors"
)

var ErrNoArgs = errors.New("no arguments provided")

// Parse コマンドライン引数を解析して設定値を取得する。
func Parse(args []string) (env.Env, port.Port, error) {
	var envZero env.Env
	var portZero port.Port

	if len(args) == 0 {
		return envZero, portZero, ErrNoArgs
	}

	fs := flag.NewFlagSet(args[0], flag.ExitOnError)
	var envString string
	{
		usage := fmt.Sprintf("environment (%s)", strings.Join(env.EnvStrings(), ","))
		fs.StringVar(&envString, "env", envZero.String(), usage)
	}
	var portString string
	{
		usage := "port number"
		fs.StringVar(&portString, "port", portZero.String(), usage)
	}
	if err := fs.Parse(args); err != nil {
		return envZero, portZero, errors.Wrap(err, 0)
	}

	envRes, err := env.EnvString(envString)
	if err != nil {
		return envZero, portZero, errors.Wrap(err, 0)
	}

	portUint, err := strconv.ParseUint(portString, 10, 16)
	if err != nil {
		return envZero, portZero, errors.Wrap(err, 0)
	}

	portRes := port.New(uint16(portUint))

	return envRes, portRes, nil
}

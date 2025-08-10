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

const flagSetName = ""

var fs = flag.NewFlagSet(flagSetName, flag.ExitOnError)

var (
	envString  string
	portString string
)

func init() {
	var envZero env.Env
	envUsage := fmt.Sprintf("environment (%s)", strings.Join(env.EnvStrings(), ","))
	fs.StringVar(&envString, "env", envZero.String(), envUsage)

	var portZero port.Port
	portUsage := "port number"
	fs.StringVar(&portString, "port", portZero.String(), portUsage)
}

func GetEnv(args []string) (env.Env, error) {
	var zero env.Env
	if err := fs.Parse(args); err != nil {
		return zero, errors.Wrap(err, 0)
	}
	if envString == "" {
		return zero, errors.New("environment is required")
	}

	e, err := env.EnvString(envString)
	if err != nil {
		return 0, err
	}

	return e, nil
}

func GetPort(args []string) (port.Port, error) {
	var zero port.Port
	if err := fs.Parse(args); err != nil {
		return zero, errors.Wrap(err, 0)
	}
	if portString == "" {
		return zero, errors.New("port is required")
	}

	i, err := strconv.ParseUint(portString, 10, 16)
	if err != nil {
		return port.Port{}, errors.Wrap(err, 0)
	}

	p := port.New(uint16(i))

	return p, nil
}

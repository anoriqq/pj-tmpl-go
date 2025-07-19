/*
Package env provides the environment domain models used throughout the application.
*/
package env

import (
	"flag"
	"log/slog"
)

//go:generate go run github.com/anoriqq/enumer@latest -type=Env -transform=lower
type Env int

var _ flag.Value = (*Env)(nil)

// Set implements flag.Value.
func (i *Env) Set(s string) error {
	slog.Info("set env", "value", s)

	e, err := EnvString(s)
	if err != nil {
		return err
	}

	*i = e

	return nil
}

const (
	_ Env = iota
	LCL
	DEV
	STG
	PRD
)

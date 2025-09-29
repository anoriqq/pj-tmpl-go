/*
Package port provides the domain layer interfaces for the application.
*/
package port

import (
	"flag"
	"fmt"
	"math"
	"strconv"

	"github.com/go-errors/errors"
)

const MaxPortValue = math.MaxUint16

var ErrInvalidPort = fmt.Errorf("port number must be between 0 and %d", MaxPortValue)

// Port ポート番号。
type Port struct {
	value uint16
}

// Set implements [flag.Value].
func (p *Port) Set(s string) error {
	v, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return errors.Wrap(err, 0)
	}

	if v > MaxPortValue {
		return errors.Wrap(ErrInvalidPort, 0)
	}
	*p = New(uint16(v))

	return nil
}

// String implements [flag.Value].
func (p *Port) String() string {
	if p.value == 0 {
		return "80"
	}

	return strconv.FormatUint(uint64(p.value), 10)
}

var _ flag.Value = (*Port)(nil)

// New ポート番号を作成する。
func New(v uint16) Port {
	v = min(v, MaxPortValue)

	return Port{value: v}
}

package port

import (
	"flag"
	"fmt"
	"math"
	"strconv"

	"github.com/go-errors/errors"
)

const maxPortValue = math.MaxUint16

var ErrInvalidPort = fmt.Errorf("port number must be between 0 and %d", maxPortValue)

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

	if v > maxPortValue {
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
	return fmt.Sprintf("%d", p.value)
}

var _ flag.Value = (*Port)(nil)

// New ポート番号を作成する。
func New(v uint16) Port {
	v = min(v, maxPortValue)

	return Port{value: v}
}

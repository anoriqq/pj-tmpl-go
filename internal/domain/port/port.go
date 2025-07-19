package port

import (
	"flag"
	"strconv"

	"github.com/go-errors/errors"
)

const maxPortValue = 65535

type Port struct {
	value uint64
}

// Set implements flag.Value.
func (p *Port) Set(s string) error {
	v, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return errors.Wrap(err, 0)
	}

	*p = New(v)

	return nil
}

// String implements flag.Value.
func (p Port) String() string {
	return strconv.FormatUint(p.value, 10)
}

var _ flag.Value = (*Port)(nil)

func New(v uint64) Port {
	v = min(v, maxPortValue)

	return Port{value: v}
}

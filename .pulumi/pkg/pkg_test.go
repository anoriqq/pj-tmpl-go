package pkg_test

import "testing"

func getOutput[T any](t *testing.T, all []any, idx int) T {
	v, ok := all[idx].(T)
	if !ok {
		t.Errorf("failed to convert type %T", *new(T))
	}

	return v
}

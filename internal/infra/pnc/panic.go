/*
Package pnc は panic に関するユーティリティを提供する。
*/
package pnc

import (
	"fmt"
	"runtime/debug"

	"github.com/go-errors/errors"
)

// Parse panicメッセージからerrorを生成する。
func Parse(msg any) error {
	stack := debug.Stack()
	text := fmt.Sprintf("panic: %v\n\n%v", msg, string(stack))

	err, parseErr := errors.ParsePanic(text)
	if parseErr != nil {
		panic(parseErr)
	}

	return err
}

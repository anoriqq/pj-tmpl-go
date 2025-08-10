package cli

import (
	"fmt"
	"runtime/debug"

	"github.com/go-errors/errors"
)

// ParsePanic panicメッセージからerrorを生成する
func ParsePanic(msg any) error {
	stack := debug.Stack()
	text := fmt.Sprintf("panic: %v\n\n%v", msg, string(stack))
	err, parseErr := errors.ParsePanic(text)
	if parseErr != nil {
		panic(parseErr)
	}

	return err
}

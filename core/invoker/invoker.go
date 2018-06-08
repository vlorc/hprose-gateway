package invoker

import (
	"context"
	"github.com/vlorc/hprose-gateway/core/types"
	"reflect"
)

type Invoker []types.Plugin

func NewInvoker(p ...types.Plugin) Invoker {
	return Invoker(p)
}

func (invoker Invoker) Invoke(ctx context.Context, name string, args []reflect.Value) ([]reflect.Value, error) {
	pos := len(invoker) - 1
	var next types.InvokeHandler
	next = func(c context.Context, m string, a []reflect.Value) (result []reflect.Value, err error) {
		if pos--; pos >= 0 {
			result, err = invoker[pos].Handler(next, c, m, a)
		}
		return
	}

	return invoker[pos].Handler(next, ctx, name, args)
}

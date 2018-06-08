package source

import (
	"context"
	"github.com/vlorc/hprose-gateway/core/types"
	"reflect"
)

type ErrorSource struct {
	err error
}

func NewErrorSource(err error) types.Source {
	return ErrorSource{err: err}
}
func (ErrorSource) SetService(*types.Service) {
}
func (ErrorSource) Service() *types.Service {
	return nil
}
func (e ErrorSource) Reset() types.Source {
	return e
}
func (e ErrorSource) Endpoint() types.Endpoint {
	return e
}
func (e ErrorSource) Invoke(context.Context, string, []reflect.Value) ([]reflect.Value, error) {
	return nil, e.err
}
func (e ErrorSource) Close() error {
	return e.err
}

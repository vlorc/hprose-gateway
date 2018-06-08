package restful

import (
	"context"
	"github.com/vlorc/hprose-gateway/core/types"
	"reflect"
)

type restfulSource struct {
	service *types.Service
}

var zero = restfulSource{}

func (h *restfulSource) SetService(service *types.Service) {
	h.service = service
}

func (h *restfulSource) Service() *types.Service {
	return h.service
}

func (p *restfulSource) Reset() types.Source {
	p.Close()
	*p = zero
	return p
}

func (h *restfulSource) Close() error {
	return nil
}

func (h *restfulSource) Endpoint() types.Endpoint {
	return h
}

func (h *restfulSource) Invoke(ctx context.Context, method string, params []reflect.Value) ([]reflect.Value, error) {
	return nil, nil
}

package forward

import (
	"context"
	"github.com/vlorc/hprose-gateway/core/types"
	"reflect"
)

type forwardSource struct {
	service *types.Service
	router  types.NamedRouter
	manager types.SourceManger
	id      string
}

var zero = forwardSource{}

func (h *forwardSource) SetService(service *types.Service) {
	if h.service = service; nil == service || len(service.Meta) <= 0 {
		return
	}
	if it, ok := service.Meta["forward.id"]; ok {
		h.id, ok = it.(string)
	} else {
		h.id = ""
	}
}

func (h *forwardSource) Service() *types.Service {
	return h.service
}

func (p *forwardSource) Reset() types.Source {
	p.Close()
	*p = zero
	return p
}

func (h *forwardSource) Close() error {
	return nil
}

func (h *forwardSource) Endpoint() types.Endpoint {
	return h
}

func (h *forwardSource) Invoke(ctx context.Context, method string, params []reflect.Value) ([]reflect.Value, error) {
	client_id := h.router.Resolver(method, h.id).Next(ctx, method, params)
	return h.manager.Resolver(client_id).Endpoint().Invoke(ctx, method, params)
}

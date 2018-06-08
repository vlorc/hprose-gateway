package hprose

import (
	"context"
	"github.com/hprose/hprose-golang/rpc"
	"github.com/vlorc/hprose-gateway/core/invoker"
	"github.com/vlorc/hprose-gateway/core/types"
	"reflect"
	"sync"
)

type hproseSource struct {
	service *types.Service
	once    sync.Once
	client  rpc.Client
	invoker invoker.Invoker
}

func (h *hproseSource) SetService(service *types.Service) {
	if nil == service {
		return
	}
	if nil != h.service && h.service.Url != service.Url && nil != h.client {
		h.client.SetURI(service.Url)
	}
	h.service = service
}

func (h *hproseSource) Service() *types.Service {
	return h.service
}

func (p *hproseSource) Reset() types.Source {
	p.Close()
	*p = zero
	return p
}

func (h *hproseSource) Close() error {
	return nil
}

func (h *hproseSource) Endpoint() types.Endpoint {
	h.once.Do(func() {
		h.client = rpc.NewClient(h.service.Url)
	})
	return h
}

func (h *hproseSource) Invoke(ctx context.Context, method string, params []reflect.Value) ([]reflect.Value, error) {
	return h.client.Invoke(method, params, settings)
}

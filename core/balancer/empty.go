package balancer

import (
	"context"
	"github.com/vlorc/hprose-gateway/core/types"
	"reflect"
)

type Factory func(map[string]struct{}, types.SourceManger) types.Balancer

func (f Factory) Update(servers map[string]struct{}, manager types.SourceManger) types.Balancer {
	return makeBalancer(servers, manager, f)
}

func (f Factory) Instance() types.Balancer {
	return f
}

func (Factory) Next(context.Context, string, []reflect.Value) string {
	return ""
}

type one struct {
	key    string
	update func(map[string]struct{}, types.SourceManger) types.Balancer
}

func (o *one) Update(servers map[string]struct{}, manager types.SourceManger) types.Balancer {
	return makeBalancer(servers, manager, o.update)
}
func (o *one) Next(context.Context, string, []reflect.Value) string {
	return o.key
}

func makeBalancer(servers map[string]struct{}, manager types.SourceManger, update func(map[string]struct{}, types.SourceManger) types.Balancer) types.Balancer {
	if 0 == len(servers) {
		return Factory(update)
	}
	if 1 == len(servers) {
		for k := range servers {
			return &one{key: k, update: update}
		}
	}
	return update(servers, manager)
}

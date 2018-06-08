package gateway

import (
	"github.com/hprose/hprose-golang/rpc"
	"github.com/vlorc/hprose-gateway/core/invoker"
	"github.com/vlorc/hprose-gateway/option"
	"reflect"
)

type HproseGateway struct {
	opt *option.GatewayOption
}

func NewGateway(o ...func(*option.GatewayOption)) *HproseGateway {
	opt := &option.GatewayOption{}
	opt.Plugins = invoker.NewInvoker(NewInvoker(opt))
	for _, v := range o {
		v(opt)
	}
	go opt.Resolver.Watch("*", opt.Water)
	return &HproseGateway{opt: opt}
}

func (g *HproseGateway) Invoke(name string, args []reflect.Value, context rpc.Context) ([]reflect.Value, error) {
	return g.opt.Plugins.Invoke(g.opt.Context, name, args)
}

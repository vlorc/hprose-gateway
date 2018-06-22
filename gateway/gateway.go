package gateway

import (
	"github.com/hprose/hprose-golang/rpc"
	"github.com/vlorc/hprose-gateway-core/invoker"
	"github.com/vlorc/hprose-gateway-core/option"
	"reflect"
	"sort"
)

type HproseGateway struct {
	opt *option.GatewayOption
}

func NewGateway(o ...func(*option.GatewayOption)) *HproseGateway {
	opt := option.NewDefault()
	opt.Plugins = invoker.NewInvoker(NewInvoker(opt))
	opt = option.NewOptionWith(opt, o...)
	sort.Sort(opt.Plugins)
	go opt.Resolver.Watch(opt.Prefix, opt.Water)
	return NewGatewayWith(opt)
}

func NewGatewayWith(opt *option.GatewayOption) *HproseGateway {
	return &HproseGateway{opt: opt}
}

func (g *HproseGateway) Invoke(name string, args []reflect.Value, context rpc.Context) ([]reflect.Value, error) {
	return g.opt.Plugins.Invoke(g.opt.Context, name, args)
}

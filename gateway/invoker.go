package gateway

import (
	"context"
	"github.com/vlorc/hprose-gateway-core/types"
	"github.com/vlorc/hprose-gateway/option"
	"go.uber.org/zap"
	"reflect"
)

type Invoker option.GatewayOption

func NewInvoker(opt *option.GatewayOption) types.Plugin {
	return (*Invoker)(opt)
}

func (g *Invoker) Close() error {
	return nil
}

func (g *Invoker) Name() string {
	return "Invoker"
}

func (g *Invoker) Handler(_ types.InvokeHandler, ctx context.Context, method string, args []reflect.Value) ([]reflect.Value, error) {
	app_id, _ := ctx.Value("appid").(string)
	client_id := g.Router.Resolver(method, app_id).Next(g.Context, method, args)
	result, err := g.Manager.Resolver(client_id).Endpoint().Invoke(g.Context, method, args)

	g.Log().Debug("Proxy",
		zap.String("path", method),
		zap.String("app_id", app_id),
		zap.String("client_id", client_id),
		zap.Error(err))
	return result, err
}

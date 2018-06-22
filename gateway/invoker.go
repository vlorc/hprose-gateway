package gateway

import (
	"context"
	"github.com/vlorc/hprose-gateway-core/option"
	"github.com/vlorc/hprose-gateway-types"
	"go.uber.org/zap"
	"reflect"
)

type Invoker option.GatewayOption

func NewInvoker(opt *option.GatewayOption) types.Plugin {
	return (*Invoker)(opt)
}

func (i *Invoker) Level() int {
	return 0
}

func (i *Invoker) Close() error {
	return nil
}

func (i *Invoker) Name() string {
	return "Invoker"
}

func (i *Invoker) Handler(_ types.InvokeHandler, ctx context.Context, method string, args []reflect.Value) ([]reflect.Value, error) {
	app_id, _ := ctx.Value("appid").(string)
	client_id := i.Router.Resolver(method, app_id).Next(i.Context, method, args)
	result, err := i.Manager.Resolver(client_id).Endpoint().Invoke(i.Context, method, args)

	i.Log().Debug("Proxy",
		zap.String("path", method),
		zap.String("app_id", app_id),
		zap.String("client_id", client_id),
		zap.Error(err))
	return result, err
}

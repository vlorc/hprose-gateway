package forward

import (
	"github.com/vlorc/hprose-gateway/core/driver"
	"github.com/vlorc/hprose-gateway/core/types"
	"sync"
)

type ForwardDriver struct {
	pool sync.Pool
}

func init() {
	driver.Register(&ForwardDriver{
		pool: sync.Pool{
			New: func() interface{} {
				return &forwardSource{}
			},
		},
	}, "forward")
}

func (f *ForwardDriver) Instance(router types.NamedRouter, manger types.SourceManger) types.Source {
	source, ok := f.pool.Get().(*forwardSource)
	if ok {
		source.router = router
		source.manager = manger
	}
	return source
}

func (f *ForwardDriver) Release(x types.Source) {
	if _, ok := x.(*forwardSource); ok {
		f.pool.Put(x.Reset())
	}
}

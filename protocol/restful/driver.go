package restful

import (
	"github.com/vlorc/hprose-gateway/core/driver"
	"github.com/vlorc/hprose-gateway/core/types"
	"sync"
)

type HproseDriver struct {
	pool sync.Pool
}

func init() {
	driver.Register(&HproseDriver{
		pool: sync.Pool{
			New: func() interface{} {
				return &restfulSource{}
			},
		},
	}, "restful")
}

func (h *HproseDriver) Instance(router types.NamedRouter, manger types.SourceManger) types.Source {
	return h.pool.Get().(types.Source)
}

func (h *HproseDriver) Release(x types.Source) {
	h.pool.Put(x.Reset())
}

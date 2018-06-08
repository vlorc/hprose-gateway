package hprose

import (
	"github.com/vlorc/hprose-gateway/core/driver"
	"github.com/vlorc/hprose-gateway/core/types"
	_ "github.com/vlorc/hprose-go-nats"
	"sync"
)

type HproseDriver struct {
	pool sync.Pool
}

func init() {
	driver.Register(&HproseDriver{
		pool: sync.Pool{
			New: func() interface{} {
				return &hproseSource{}
			},
		},
	}, "hprose", "")
}

func (h *HproseDriver) Instance(types.NamedRouter, types.SourceManger) types.Source {
	return h.pool.Get().(types.Source)
}

func (h *HproseDriver) Release(x types.Source) {
	h.pool.Put(x.Reset())
}

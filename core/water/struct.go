package water

import (
	"github.com/vlorc/hprose-gateway/core/types"
	"sync"
)

type ChannelWater chan []types.Update

type SnapshotWater struct {
	types.NamedWatcher
	water   []types.NamedWatcher
	router  types.NamedRouter
	manager types.SourceManger
	service sync.Map
	channel chan types.Update
	push    func(types.Update)
}

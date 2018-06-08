package driver

import "github.com/vlorc/hprose-gateway/core/types"

var driver = make(map[string]types.SourceFactory)

func Query(name string) types.SourceFactory {
	if factory, ok := driver[name]; ok {
		return factory
	}
	return nil
}

func Register(factory types.SourceFactory, name ...string) {
	for _, v := range name {
		driver[v] = factory
	}
}

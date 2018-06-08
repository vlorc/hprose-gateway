package plugin

import "github.com/vlorc/hprose-gateway/core/types"

var driver = make(map[string]types.PluginFactory)

func Query(name string) types.PluginFactory {
	if factory, ok := driver[name]; ok {
		return factory
	}
	return nil
}

func Register(factory types.PluginFactory, name ...string) {
	for _, v := range name {
		driver[v] = factory
	}
}

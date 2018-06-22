package named

import (
	"github.com/vlorc/hprose-gateway-types"
	"strings"
)

type ModuleNamed struct{}
type AppNamed struct{}

func (ModuleNamed) Resolver(path string, _ string) interface{} {
	return rootModule(path)
}

func rootModule(path string) string {
	if pos := strings.IndexByte(path, byte('_')); pos > 0 {
		path = path[:pos]
	} else {
		path = ""
	}
	return path
}

func (AppNamed) Resolver(path string, id string) interface{} {
	return types.Named{
		Path: rootModule(path),
		Id:   id,
	}
}

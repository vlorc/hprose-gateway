package session

import "github.com/vlorc/hprose-gateway/core/plugin"

func init() {
	plugin.Register(sessionParamFactory{}, "session")
}

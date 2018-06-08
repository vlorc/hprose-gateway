package balancer

import (
	"context"
	"github.com/vlorc/hprose-gateway/core/types"
	mrand "math/rand"
	"reflect"
	"time"
)

var rand *mrand.Rand = mrand.New(mrand.NewSource(time.Now().UnixNano()))

type random []string

var RandomFactory Factory = func(servers map[string]struct{}, _ types.SourceManger) types.Balancer {
	ss := make([]string, 0, len(servers))
	for k := range servers {
		ss = append(ss, k)
	}
	return random(ss)
}

func (random) Update(servers map[string]struct{}, manager types.SourceManger) types.Balancer {
	return RandomFactory.Update(servers, manager)
}

func (r random) Next(context.Context, string, []reflect.Value) string {
	i := rand.Uint32() % uint32(len(r))
	return r[i]
}

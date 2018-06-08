package balancer

import (
	"context"
	"encoding/json"
	"github.com/vlorc/hprose-gateway/core/types"
	"reflect"
	"strconv"
)

type __weight struct {
	current int
	value   int
	key     string
}

type weight struct {
	service []__weight
	value   int
	update  Factory
}

func WeightFactoryWith(format func(string, types.SourceManger) int) (update Factory) {
	update = func(servers map[string]struct{}, manager types.SourceManger) types.Balancer {
		ss := make([]__weight, 0, len(servers))
		val := 0
		for k, _ := range servers {
			v := format(k, manager)
			val += v
			ss = append(ss, __weight{key: k, value: v})
		}
		return &weight{
			service: ss,
			value:   val,
			update:  update,
		}
	}
	return
}

func (w *weight) Update(servers map[string]struct{}, manager types.SourceManger) types.Balancer {
	return w.update.Update(servers, manager)
}

func (w *weight) Next(context.Context, string, []reflect.Value) string {
	index := 0
	value := w.service[0].current
	for i := range w.service {
		if w.service[i].current += w.service[i].value; w.service[i].current > value {
			value = w.service[i].current
			index = i
		}
	}
	w.service[index].current -= w.value
	return w.service[index].key
}

func formatInt(key string) func(string, types.SourceManger) int {
	return func(id string, manager types.SourceManger) (val int) {
		val = 1
		source := manager.Resolver(id)
		if nil == source {
			return
		}
		service := source.Service()
		if nil == service || len(service.Meta) <= 0 {
			return
		}
		if val = safeInteger(service.Meta[key], 1); 0 == val {
			val = 1
		} else if val < 0 {
			val = -val
		}
		return
	}
}

func safeInteger(s interface{}, v int) int {
	if nil == s {
		return v
	}
	switch r := s.(type) {
	case string:
		if i, err := strconv.ParseInt(r, 10, 64); nil == err {
			v = int(i)
		}
	case json.Number:
		if i, err := r.Int64(); nil == err {
			v = int(i)
		}
	case []byte:
		if i, err := strconv.ParseInt(string(r), 10, 64); nil == err {
			v = int(i)
		}
	case float64:
		v = int(r)
	case float32:
		v = int(r)
	case uint:
		v = int(r)
	case int:
		v = int(r)
	case uint64:
		v = int(r)
	case int64:
		v = int(r)
	case uint32:
		v = int(r)
	case int32:
		v = int(r)
	case uint16:
		v = int(r)
	case int16:
		v = int(r)
	case byte:
		v = int(r)
	}
	return v
}

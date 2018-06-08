package hprose

import (
	"github.com/hprose/hprose-golang/rpc"
	"reflect"
)

var settings = &rpc.InvokeSettings{
	Mode:        rpc.Raw,
	ResultTypes: []reflect.Type{reflect.TypeOf(([]byte)(nil))},
}

var zero = hproseSource{}

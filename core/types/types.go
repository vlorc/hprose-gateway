package types

import (
	"context"
	"reflect"
	"sync"
)

type Operation uint8

const (
	Add Operation = iota
	Delete
)

func (o Operation) String() string {
	return []string{"PUT", "DELETE"}[o]
}

type Source interface {
	SetService(*Service)
	Service() *Service
	Reset() Source
	Endpoint() Endpoint
	Close() error
}

type Describe struct {
	Name  string            `json:",omitempty"`
	Param map[string]string `json:",omitempty"`
}

type Service struct {
	Id       string                 `json:",omitempty"`
	Name     string                 `json:",omitempty"`
	Path     string                 `json:",omitempty"`
	Driver   string                 `json:",omitempty"`
	Version  string                 `json:",omitempty"`
	Url      string                 `json:",omitempty"`
	Platform string                 `json:",omitempty"`
	Plugins  []Describe             `json:",omitempty"`
	Meta     map[string]interface{} `json:",omitempty"`
}

type Named struct {
	Id   string
	Path string
}

type Update struct {
	Op      Operation
	Id      string
	Service *Service `json:",omitempty"`
}

type NamedNode struct {
	sync.Locker
	Peer     map[string]struct{}
	Balancer Balancer
}

type NamedRegister interface {
	Update(service *Service) error
	Close() error
}

type NamedManger interface {
	Register(name, uuid string) NamedRegister
	Keys() []string
	Close() error
}

type NamedResolver interface {
	All(name string, watcher NamedWatcher) error
	Watch(name string, watcher NamedWatcher) error
	Close() error
}

type NamedWatcher interface {
	Push([]Update) error
	Pop() ([]Update, error)
	Close() error
}

type SourceManger interface {
	Resolver(string) Source
	Append(string, Source) SourceManger
	Remove(string) SourceManger
	Close() error
}

type Endpoint interface {
	Invoke(context.Context, string, []reflect.Value) ([]reflect.Value, error)
	Close() error
}

type Balancer interface {
	Update(map[string]struct{}, SourceManger) Balancer
	Next(context.Context, string, []reflect.Value) string
}

type NamedMode interface {
	Resolver(path string, id string) interface{}
}

type NamedRouter interface {
	Resolver(path string, id string) Balancer
	Append(name Named, id string) Balancer
	Remove(name Named, id string) Balancer
}

type BalancerFactory interface {
	Instance() Balancer
}

type EndpointFactory interface {
	Instance(url string) Endpoint
}

type PluginFactory interface {
	Instance(ctx context.Context, param map[string]string) Plugin
}

type InvokeHandler func(context.Context, string, []reflect.Value) ([]reflect.Value, error)

type Plugin interface {
	Name() string
	Handler(InvokeHandler, context.Context, string, []reflect.Value) ([]reflect.Value, error)
	Close() error
}

type SourceFactory interface {
	Instance(router NamedRouter, manger SourceManger) Source
	Release(Source)
}

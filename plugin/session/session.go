package session

import (
	"context"
	"errors"
	"github.com/vlorc/hprose-gateway/core/types"
	"reflect"
	"regexp"
	"strings"
)

type sessionParamFactory struct{}

type sessionParam struct {
	factory SessionFactory
	err     error
	ignore  func(string) bool
	id      string
}

func ignore(mode, match string) (result func(string) bool) {
	switch mode {
	case "prefix":
		result = func(s string) bool {
			return strings.HasPrefix(s, match)
		}
	case "suffix":
		result = func(s string) bool {
			return strings.HasSuffix(s, match)
		}
	case "find":
		result = func(s string) bool {
			return strings.Index(s, match) >= 0
		}
	case "regexp":
		result = regexp.MustCompile(match).MatchString
	default:
		result = func(string) bool {
			return false
		}
	}
	return
}

func (sessionParamFactory) Instance(ctx context.Context, param map[string]string) types.Plugin {
	factory := ctx.Value("SessionFactory").(func(string) SessionFactory)
	return &sessionParam{
		factory: factory(param["secret"]),
		err:     errors.New(param["error"]),
		ignore:  ignore(param["ignore.mode"], param["ignore.match"]),
		id:      param["appid"],
	}
}

func (s *sessionParam) Close() error {
	return nil
}

func (s *sessionParam) Name() string {
	return "session"
}

func (s *sessionParam) Handler(next types.InvokeHandler, ctx context.Context, method string, args []reflect.Value) ([]reflect.Value, error) {
	if len(args) <= 0 {
		return nil, s.err
	}
	if s.ignore(method) {
		return next(context.WithValue(ctx, "appid", s.id), method, args)
	}
	str, ok := args[len(args)-1].Interface().(string)
	if !ok {
		return nil, s.err
	}
	session, err := s.factory.Instance(str)
	if nil != err || nil != session.Verify() {
		return nil, s.err
	}
	ctx = context.WithValue(ctx, "appid", session.AppId())
	args[len(args)-1] = reflect.ValueOf(session.Value())
	return next(ctx, method, args)
}

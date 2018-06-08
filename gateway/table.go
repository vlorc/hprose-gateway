package gateway

import (
	"bytes"
	"fmt"
	"github.com/vlorc/hprose-gateway/core/types"
	"io"
	"reflect"
	"strconv"
	"sync"
)

type HproseTable struct {
	method map[string]struct{}
	buffer bytes.Buffer
	head   []byte
	lock   sync.RWMutex
	key    string
}

func NewTable(key string) *HproseTable {
	return &HproseTable{
		method: make(map[string]struct{}),
		head:   []byte("Fa2{u#u*"),
		key:    key,
	}
}

func (t *HproseTable) Push(updates []types.Update) error {
	for i := range updates {
		if types.Add == updates[i].Op && nil != updates[i].Service && len(updates[i].Service.Meta) > 0 {
			t.append(updates[i].Service.Meta[t.key])
		}
	}
	return nil
}

func (t *HproseTable) Pop() ([]types.Update, error) {
	return nil, nil
}

func (t *HproseTable) Close() error {
	return nil
}

func (t *HproseTable) writeName(val interface{}) bool {
	name := fmt.Sprint(val)
	if "#" == name || "*" == name {
		return false
	}
	if _, ok := t.method[name]; !ok {
		t.method[name] = struct{}{}
		t.buffer.WriteString("s")
		t.buffer.WriteString(strconv.Itoa(len(name)))
		t.buffer.WriteString(`"`)
		t.buffer.WriteString(name)
		t.buffer.WriteString(`"`)
		return true
	}
	return false
}
func (t *HproseTable) append(table interface{}) {
	val := reflect.ValueOf(table)
	if !val.IsValid() || val.IsNil() || (reflect.Slice != val.Kind() && reflect.Array != val.Kind()) && val.Len() <= 0 {
		return
	}

	t.lock.Lock()
	defer t.lock.Unlock()

	count := 0
	for i, l := 0, val.Len(); i < l; i++ {
		if t.writeName(val.Index(i).Interface()) {
			count++
		}
	}
	if count > 0 {
		t.head = []byte(fmt.Sprintf("Fa%d{u#u*", len(t.method)+2))
	}
}

func (t *HproseTable) Pipe(w io.Writer) {
	t.lock.RLock()
	defer t.lock.RUnlock()

	w.Write(t.head)
	w.Write(t.buffer.Bytes())
	w.Write([]byte("}z"))
}

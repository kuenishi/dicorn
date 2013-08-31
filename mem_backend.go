package dicorn

import (
	"fmt"
	"strconv"
	"sync"
)

type MemoryBackend struct {
	data  map[string][]byte
	mutex sync.Mutex
}

type NotFoundError struct {
}

func (e *NotFoundError) Error() string {
	return "not_found"
}
func NewMemoryBackend() *MemoryBackend {
	m := &MemoryBackend{}
	m.Init()
	return m
}
func (mb *MemoryBackend) Init() {
	mb.data = make(map[string][]byte)
}
func (mb *MemoryBackend) Set(key string, flag, expire, size int, value []byte) error {
	if size == len(value) {
		mb.mutex.Lock()
		mb.data[key] = value
		mb.mutex.Unlock()
		// TODO: how can I make this constant?
		return nil
	} // else
	return &NotFoundError{}
}
func (mb *MemoryBackend) Get(key string) ([]byte, error) {
	mb.mutex.Lock()
	v, ok := mb.data[key]
	mb.mutex.Unlock()
	if ok {
		return v, nil
	}
	return nil, &NotFoundError{}
}
func (mb *MemoryBackend) Delete(key string) error {
	mb.mutex.Lock()
	_, ok := mb.data[key]
	if ok {
		delete(mb.data, key)
		mb.mutex.Unlock()
		return nil
	}
	mb.mutex.Unlock() // better use defer
	return &NotFoundError{}
}
func (mb *MemoryBackend) Incr(key string, value int) (int, error) {
	mb.mutex.Lock()
	v, ok := mb.data[key]
	if ok {
		v0, e0 := strconv.Atoi(string(v))
		nv := v0 + value
		if e0 == nil {
			str := fmt.Sprintf("%d", nv)
			mb.data[key] = []byte(str)
			mb.mutex.Unlock()
			return nv, nil
		} else {
			mb.mutex.Unlock()
			return 0, nil
		}
	}
	mb.mutex.Unlock()
	return 0, &NotFoundError{}
}
func (mb *MemoryBackend) Decr(key string, value int) (int, error) {
	mb.mutex.Lock()
	v, ok := mb.data[key]
	if ok {
		v0, e0 := strconv.Atoi(string(v))
		nv := max(v0-value, 0)
		if e0 == nil {
			str := fmt.Sprintf("%d", nv)
			mb.data[key] = []byte(str)
			mb.mutex.Unlock()
			return nv, nil
		} else {
			mb.mutex.Unlock()
			return 0, nil
		}
	}
	mb.mutex.Unlock()
	return 0, &NotFoundError{}
}

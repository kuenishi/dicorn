package dicorn

import (
	"strings"
	// "github.com/mrb/riakpbc"
)

type RiakBackend struct {
	hosts []string
}

func NewRiakBackend(hosts string) *RiakBackend {
	h := strings.Split(hosts, ",")
	// TODO: check riak existence here
	return &RiakBackend{hosts: h}
}

func (mb *RiakBackend) Init() {
}

func (mb *RiakBackend) Set(key string, flag, expire, size int, value []byte) error {
	return &NotFoundError{}
}
func (mb *RiakBackend) Get(key string) ([]byte, error) {
	return nil, &NotFoundError{}
}
func (mb *RiakBackend) Delete(key string) error {
	return &NotFoundError{}
}
func (mb *RiakBackend) Incr(key string, value int) (int, error) {
	return 0, &NotFoundError{}
}
func (mb *RiakBackend) Decr(key string, value int) (int, error) {
	return 0, &NotFoundError{}
}

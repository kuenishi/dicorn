package dicorn

import (
	"strings"
	"github.com/mrb/riakpbc"
)

const BUCKET = "fake_memcached"

type RiakBackend struct {
	hosts []string
}

func NewRiakBackend(hosts string) *RiakBackend {
	h := strings.Split(hosts, ",")
	// TODO: check riak existence here
	return &RiakBackend{hosts: h}
}

func (rb *RiakBackend) Init() {
}

func (rb *RiakBackend) Set(key string, flag, expire, size int, value []byte) error {
	c := rb.connect()
	_, err := c.StoreObject(BUCKET, key, value)
	c.Close()
	return err
}
func (rb *RiakBackend) Get(key string) ([]byte, error) {
	c := rb.connect()
	obj, err := c.FetchObject(BUCKET, key)
	c.Close()
	if err != nil {
		return nil, &NotFoundError{}
	}
	return (obj.GetContent()[0].GetValue()), nil
}
func (rb *RiakBackend) Delete(key string) error {
	return &NotFoundError{}
}
func (rb *RiakBackend) Incr(key string, value int) (int, error) {
	return 0, &NotFoundError{}
}
func (rb *RiakBackend) Decr(key string, value int) (int, error) {
	return 0, &NotFoundError{}
}

func (rb *RiakBackend) connect() *riakpbc.Client {
	return riakpbc.NewClient(rb.hosts)
}
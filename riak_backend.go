package dicorn

import (
	"fmt"
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
	c := riakpbc.NewClient(h)
	if c.Dial() == nil {
		fmt.Printf("riak is ok: %v\n", c)
	}
	c.Close()
	return &RiakBackend{hosts: h}
}

func (rb *RiakBackend) Init() {
}

func (rb *RiakBackend) Set(key string, flag, expire, size int, value []byte) error {
	c := rb.connect()
	//fmt.Printf("%s %s %v %v\n", BUCKET, key, string(value), c)
	_, err := c.StoreObject(BUCKET, key, value)
	//fmt.Printf(" => %v / %v\n", r, err)
	rb.disconnect(c)
	return err
}
func (rb *RiakBackend) Get(key string) ([]byte, error) {
	c := rb.connect()
	obj, err := c.FetchObject(BUCKET, key)
	rb.disconnect(c)
	//fmt.Printf(" => %v / %v\n", obj, err)
	if err != nil {
		return nil, &NotFoundError{}
	}
	return (obj.GetContent()[0].GetValue()), nil
}
func (rb *RiakBackend) Delete(key string) error {
	c := rb.connect()
	_, err := c.DeleteObject(BUCKET, key)
	rb.disconnect(c)
	return err
}
func (rb *RiakBackend) Incr(key string, value int) (int, error) {
	return 0, &NotFoundError{}
}
func (rb *RiakBackend) Decr(key string, value int) (int, error) {
	return 0, &NotFoundError{}
}

func (rb *RiakBackend) connect() *riakpbc.Client {
	c := riakpbc.NewClient(rb.hosts)
	c.Dial()
	return c
}
func (rb *RiakBackend) disconnect(c *riakpbc.Client) {
	c.Close()
}
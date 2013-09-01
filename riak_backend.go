package dicorn

import (
	"container/list"
	"github.com/mrb/riakpbc"
	"log"
	"strings"
	"sync"
)

const BUCKET = "fake_memcached"

type RiakBackend struct {
	hosts []string
	mutex sync.Mutex
	list  list.List
}

func NewRiakBackend(hosts string) *RiakBackend {
	h := strings.Split(hosts, ",")
	// TODO: check riak existence here
	rb := &RiakBackend{hosts: h}
	rb.Init()
	return rb
}

func (rb *RiakBackend) Init() {
	rb.list.Init()
	c := riakpbc.NewClient(rb.hosts)
	if c.Dial() == nil {
		log.Println("riak is ok:", c)
		rb.mutex.Lock()
		rb.list.PushBack(c)
		rb.mutex.Unlock()
	} else {
		log.Println("warning: cannot connect to riak")
	}
}

func (rb *RiakBackend) Set(key string, flag, expire, size int, value []byte) error {
	c := rb.connect()
	//fmt.Printf("%s %s %v %v\n", BUCKET, key, string(value), c)
	r, err := c.StoreObject(BUCKET, key, value)
	if err != nil {
		c.Close()
		log.Printf(" => %v / %v\n", r, err.Error())
	}
	rb.disconnect(c)
	return err
}
func (rb *RiakBackend) Get(key string) ([]byte, error) {
	c := rb.connect()
	obj, err := c.FetchObject(BUCKET, key)
	if err != nil {
		c.Close()
		log.Printf(" => %v / %v\n", obj, err.Error())
		return nil, &NotFoundError{}
	}
	rb.disconnect(c)
	return (obj.GetContent()[0].GetValue()), nil
}
func (rb *RiakBackend) Delete(key string) error {
	c := rb.connect()
	r, err := c.DeleteObject(BUCKET, key)
	if err != nil {
		c.Close()
		log.Printf(" => %v / %v\n", r, err.Error())
		return err
	}
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
	rb.mutex.Lock()
	if rb.list.Len() == 0 {
		rb.mutex.Unlock()
		c := riakpbc.NewClient(rb.hosts)
		c.Dial()
		log.Println("new connection acquired!!!")
		return c
	}
	c := rb.list.Front()
	rb.list.Remove(c)
	rb.mutex.Unlock()
	return c.Value.(*riakpbc.Client)
}
func (rb *RiakBackend) disconnect(c *riakpbc.Client) {
	rb.mutex.Lock()
	rb.list.PushBack(c)
	rb.mutex.Unlock()
}

package dicorn

type RiakBackend struct{}

func NewRiakBackend(hosts string) *RiakBackend {
	return nil
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

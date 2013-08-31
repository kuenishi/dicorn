package dicorn

type Backend interface {
	// so far interfaces for Memcached API
	Init()
	Set(string, int, int, int, []byte) error
	Get(key string) ([]byte, error)
	Delete(key string) error
	Incr(key string, valur int) (int, error)
	Decr(key string, valur int) (int, error)
}

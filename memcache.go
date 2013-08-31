package dicorn

// memcache protocol implementation
// https://github.com/memcached/memcached/blob/master/doc/protocol.txt

import (
	"bytes"
	"fmt"
	"net"
	"strconv"
)

type MemcacheAPI struct {
	Backend
}

//	HandleBytes(buf []byte, conn net.Conn)

func NewMemcacheAPI(backend, backend_hosts string) *MemcacheAPI {
	switch backend {
	case "mem":
		m := &MemcacheAPI{Backend: &MemoryBackend{}}
		m.Backend.Init()
		return m
	case "riak":
		r := NewRiakBackend(backend_hosts)
		m := &MemcacheAPI{Backend: r}
		return m
	}
	return nil
}

func (m *MemcacheAPI) set(key string, flag, expire, size int, value []byte, conn net.Conn) {
	err := m.Backend.Set(key, flag, expire, size, value)
	if err == nil {
		conn.Write([]byte("STORED\r\n"))
	} else {
		conn.Write([]byte("ERROR\r\n"))
	}
}
func (m *MemcacheAPI) get(key string, conn net.Conn) {
	v, err := m.Backend.Get(key)
	if err == nil {
		ret_str := fmt.Sprintf("VALUE %s 0 %d\r\n", key, len(v))
		conn.Write([]byte(ret_str))
		conn.Write(v)
		conn.Write([]byte("\r\n"))
		conn.Write([]byte("END\r\n"))
	} else {
		conn.Write([]byte("NOT_FOUND\r\n"))
	}
}

func (m *MemcacheAPI) delete(key string, conn net.Conn) {
	err := m.Backend.Delete(key)
	if err == nil {
		conn.Write([]byte("DELETED\r\n"))
	}
	conn.Write([]byte("NOT_FOUNDr\n"))
}

func (m *MemcacheAPI) incr(key, value string, conn net.Conn) {
	v, e := strconv.Atoi(value)
	if e != nil {
		conn.Write([]byte("BAD_VALUE\r\n"))
		return
	}
	i, err := m.Backend.Incr(key, v)
	if err != nil {
		conn.Write([]byte("NOT_FOUND\r\n"))
	}
	ret_str := fmt.Sprintf("%d\r\n", i)
	conn.Write([]byte(ret_str))
}

func max(lhs, rhs int) int {
	ret := rhs
	if lhs >= rhs {
		return lhs
	}
	return ret
}

func (m *MemcacheAPI) decr(key, value string, conn net.Conn) {
	v, e := strconv.Atoi(value)
	if e != nil {
		conn.Write([]byte("BAD_VALUE\r\n"))
		return
	}
	i, err := m.Backend.Decr(key, v)
	if err != nil {
		conn.Write([]byte("BAD_VALUE\r\n"))
	}
	ret_str := fmt.Sprintf("%d\r\n", i)
	conn.Write([]byte(ret_str))
}

func (m *MemcacheAPI) HandleBytes(buf []byte, conn net.Conn) {

	lines := bytes.Split(buf, []byte("\r\n"))
	//fmt.Printf("> %#v\n", string(buf))

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		tokens := bytes.Split(line, []byte(" "))
		cmd := string(tokens[0])
		//fmt.Printf("%d> %#v <- %s\n", i, string(line), cmd)
		switch cmd {
		case "":
			continue
		//Storage commands
		// <command name> <key> <flags> <exptime> <bytes> [noreply]\r\n
		case "set":
			v := lines[i+1]
			k := string(tokens[1])
			f, _ := strconv.Atoi(string(tokens[2]))
			e, _ := strconv.Atoi(string(tokens[3]))
			b, _ := strconv.Atoi(string(tokens[4]))
			m.set(k, f, e, b, v, conn)
			i = i + 1

		case "add":
			notSupported(conn)
			i = i + 1
		case "replace":
			notSupported(conn)
			i = i + 1
		case "append":
			notSupported(conn)
			i = i + 1
		case "prepend":
			notSupported(conn)
			i = i + 1
		case "cas":
			notSupported(conn)
			i = i + 1

			//Retrieval command
		case "get":
			k := string(tokens[1])
			m.get(k, conn)

		//case "gets":
		case "delete":
			k := string(tokens[1])
			m.delete(k, conn)

		case "incr":
			k := string(tokens[1])
			v := string(tokens[2])
			m.incr(k, v, conn)

		case "decr":
			k := string(tokens[1])
			v := string(tokens[2])
			m.decr(k, v, conn)

		//case "touch":
		//case "slabs":
		//case "stats":
		//case "flush_all":
		//case "version":
		//case "quit":
		default:
			//conn.Write([]byte("ERROR\r\n"))
			//continue looping
		}
	}
	// conn.Close()
}

func notSupported(conn net.Conn) {
	conn.Write([]byte("SERVER_ERROR not implemented\r\n"))
}

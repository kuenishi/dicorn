package dicorn
// memcache protocol implementation
// https://github.com/memcached/memcached/blob/master/doc/protocol.txt

import (
	"net"
	"fmt"
	"bytes"
	"sync"
	"strconv"
	)


type MemcacheAPI struct {
	data map[string][]byte
	mutex sync.Mutex
}
//	HandleBytes(buf []byte, conn net.Conn)

func NewMemcacheAPI() *MemcacheAPI {
	ret := &MemcacheAPI{}
	ret.data = make(map[string][]byte)
	return ret
}

func (m *MemcacheAPI) set(key string, flag, expire, size int, value []byte, conn net.Conn) {
	if size == len(value) {
		m.mutex.Lock()
		m.data[key] = value
		m.mutex.Unlock()
		// TODO: how can I make this constant?
		conn.Write([]byte("STORED\r\n"))
	} // else 
}

func (m *MemcacheAPI) HandleBytes(buf []byte, conn net.Conn) {

	lines := bytes.Split(buf, []byte("\r\n"))
	//fmt.Printf("> %#v\n", lines)

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		tokens := bytes.Split(line, []byte(" "))
		cmd := string(tokens[0])
		//fmt.Printf("%d> %#v\n", i, string(line))
		switch cmd {

			//Storage commands
			// <command name> <key> <flags> <exptime> <bytes> [noreply]\r\n
		case "set":
			v := lines[i+1]
			k := string(tokens[1])
			f, _ := strconv.Atoi(string(tokens[2]))
			e, _ := strconv.Atoi(string(tokens[3]))
			b, _ := strconv.Atoi(string(tokens[4]))
			m.set(k, f, e, b, v, conn)
			i = i+1

		case "add":
			notSupported(conn)
			i = i+1
		case "replace":
			notSupported(conn)
			i = i+1
		case "append":
			notSupported(conn)
			i = i+1
		case "prepend":
			notSupported(conn)
			i = i+1
		case "cas":
			notSupported(conn)
			i = i+1

			//Retrieval command
		case "get":
			k := string(tokens[1])
			m.mutex.Lock()
			v := m.data[k]
			m.mutex.Unlock()
			ret_str := fmt.Sprintf("VALUE %s 0 %d\r\n", k, len(v))
			conn.Write([]byte(ret_str))
			conn.Write(v)
			conn.Write([]byte("\r\n"))

		//case "gets":
		//case "delete":
		//case "incr":
		//case "decr":
		//case "touch":
		//case "slabs":
		//case "stats":
		//case "flush_all":
		//case "version":
		//case "quit":
		default:
			conn.Write([]byte("ERROR\r\n"))
		}
	}
	conn.Close()
}

func notSupported(conn net.Conn) {
	conn.Write([]byte("SERVER_ERROR not implemented\r\n"))
}
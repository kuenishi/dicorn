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
func (m *MemcacheAPI) get(key string, conn net.Conn) {
	m.mutex.Lock()
	v := m.data[key]
	m.mutex.Unlock()
	ret_str := fmt.Sprintf("VALUE %s 0 %d\r\n", key, len(v))
	conn.Write([]byte(ret_str))
	conn.Write(v)
	conn.Write([]byte("\r\n"))
}

func (m *MemcacheAPI) delete(key string, conn net.Conn) {
	m.mutex.Lock()
	_, ok := m.data[key]
	str := ""
	if ok {
		delete(m.data, key)
		str = "DELETED"
	} else {
		str = "NOT_FOUND"
	}
	m.mutex.Unlock()
	ret_str := fmt.Sprintf("%s\r\n", str)
	conn.Write([]byte(ret_str))
}

func (m *MemcacheAPI) incr(key, value string, conn net.Conn) {
	m.mutex.Lock()
	v, ok := m.data[key]
	str := ""
	if ok {
		v0, e0 := strconv.Atoi(string(v))
		v1, e1 := strconv.Atoi(value)
		if e0 == nil && e1 == nil {
			str = fmt.Sprintf("%d", v0 + v1)
			m.data[key] = []byte(str)
		} else {
			str = "ERROR?"
		}
	} else {
		str = "NOT_FOUND"
	}
	m.mutex.Unlock()
	ret_str := fmt.Sprintf("%s\r\n", str)
	conn.Write([]byte(ret_str))
}

func max(lhs, rhs int) (int) {
	ret := rhs
	if lhs >= rhs {
		return lhs
	}
	return ret
}

func (m *MemcacheAPI) decr(key, value string, conn net.Conn) {
	m.mutex.Lock()
	v, ok := m.data[key]
	str := ""
	if ok {
		v0, e0 := strconv.Atoi(string(v))
		v1, e1 := strconv.Atoi(value)
		if e0 == nil && e1 == nil {
			str = fmt.Sprintf("%d", max(v0 - v1, 0))
			m.data[key] = []byte(str)
		} else {
			str = "ERROR"
		}
	} else {
		str = "NOT_FOUND"
	}
	m.mutex.Unlock()
	ret_str := fmt.Sprintf("%s\r\n", str)
	conn.Write([]byte(ret_str))
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
			conn.Write([]byte("ERROR\r\n"))
		}
	}
	conn.Close()
}

func notSupported(conn net.Conn) {
	conn.Write([]byte("SERVER_ERROR not implemented\r\n"))
}
package dicorn

import (
	"io"
	"log"
	"net"
	"os"
)

type StorageInterface interface {
	HandleBytes(buf []byte, conn net.Conn)
}

func Version() string {
	return "dicorn version -200M"
}

func Run(addr, backend_type, backend_hosts string) {
	//log.Printf("listening on %s (accessing riak: %s)\n", addr, riak_addr)
	log.Printf("listening on %s\n", addr)
	log.Printf("backend: %s\n", backend_type)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		println("error")
		os.Exit(-1)
	}

	var api StorageInterface
	api = NewMemcacheAPI(backend_type, backend_hosts)
	for {
		conn, err := listener.Accept()
		if err != nil {
			println("fail on accept")
			return
		}
		go handleAccept(conn, api)
	}
}

func handleAccept(conn net.Conn, api StorageInterface) {
	defer func() {
		conn.Close()
		// log.Printf("%s connection from %s closed\n",
		// 	conn.RemoteAddr().Network(),
		// 	conn.RemoteAddr().String())
	}()
	buf := make([]byte, 65536)
	for {
		n, err := conn.Read(buf)
		if err == io.EOF {
			log.Printf("connection closed: %s\n", err.Error())
			return
		} else if neterr, ok := err.(net.Error); ok && neterr.Timeout() {
			log.Printf("timeout: %s\n", err.Error())
			return
		} else if err != nil {
			// log.Printf("fail on reading, or probably disconneted: %s\n", err.Error())
			log.Printf(err.Error())
			return
		}
		//log.Printf("%s connection from %s, recv'd %d bytes\n",
		//conn.RemoteAddr().Network(), conn.RemoteAddr().String(), n)

		// TODO: handlebytes cannot handle boundary well,
		// this function currently assumes this buf[:n]
		// with ending in message boundary.
		api.HandleBytes(buf[:n], conn)
	}
}

package dicorn

import (
	"fmt"
	"os"
	"net"
	)

type StorageInterface interface {
	HandleBytes(buf []byte, conn net.Conn)
}

func Run(addr string){
	fmt.Printf("listening on %s\n", addr)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		println("error")
		os.Exit(-1)
	}
	
	var api StorageInterface
	api = NewMemcacheAPI()
	for {
		conn, err := listener.Accept()
		if err != nil {
			println("fail on accept")
			return
		}
		go handleAccept(conn, api)
	}
}

func handleAccept(conn net.Conn, api StorageInterface){
	defer func() {
		conn.Close()
		// fmt.Printf("connection from %v closed\n", conn)
	}()
	buf := make([]byte, 65536)
	n, err := conn.Read(buf)
	if err != nil {
		println("fail on reading")
		return
	}
	fmt.Printf("%s connection from %s, recv'd %d bytes\n",
		conn.RemoteAddr().Network(), conn.RemoteAddr().String(), n)

	api.HandleBytes(buf[:n], conn)
}
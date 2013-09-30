package main

import (
	"fmt"
	"os"
	"flag"
	"dicorn"
	"log"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: dicorn\n")
	flag.PrintDefaults()
	os.Exit(1)
}

func main() {

	flag.Usage = usage
	//front
	subcmd := flag.String("cmd", "help", "dicorn type [memcache|cassandra|...]")
	listen := flag.String("listen", "localhost:9979", "listen port")
	//backend
	backend := flag.String("backend", "mem", "[mem|riak]")
	backend_hosts := flag.String("host", "localhost:8087", "riak pb ports")
	flag.Parse()
	// fmt.Printf("%v %v %v\n", *subcmd, *listen, *riak_hosts)

	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)
	switch *subcmd {
	case "memcache":
		dicorn.Run(*listen, *backend, *backend_hosts)
	case "cassandra":
		dicorn.RunCass(*listen, *backend, *backend_hosts)
	case "help": usage()
	default:     usage()
	}

}

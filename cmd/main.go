package main

import (
	"fmt"
	"os"
	"flag"
	"dicorn"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: dicorn\n")
	flag.PrintDefaults()
	os.Exit(1)
}

func main() {
	//os.Getenv("HOME") + "/.s3cfg")

	flag.Usage = usage
	//front
	subcmd := flag.String("cmd", "help", "dicorn type [memcache|...]")
	listen := flag.String("listen", "localhost:9979", "listen port")
	//backend
	backend := flag.String("backend", "mem", "[mem|riak]")
	backend_hosts := flag.String("host", "localhost:8087", "riak pb ports")
	flag.Parse()
	// fmt.Printf("%v %v %v\n", *subcmd, *listen, *riak_hosts)

	switch *subcmd {
	case "memcache":
		dicorn.Run(*listen, *backend, *backend_hosts)
	case "help": usage()
	default:     usage()
	}

}

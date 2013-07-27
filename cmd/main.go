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
	subcmd := flag.String("cmd", "help", "dicorn type [memcache|...]")
	listen := flag.String("listen", "localhost:9979", "listen port")
	//riak_hosts := flag.String("riak", "localhost:8087", "riak pb ports")
	flag.Parse()
	// fmt.Printf("%v %v %v\n", *subcmd, *listen, *riak_hosts)

	switch *subcmd {
	case "memcache":
		dicorn.Run(*listen)

	case "help": usage()
	default:     usage()
	}

}

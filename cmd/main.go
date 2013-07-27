package main

import (
	"fmt"
	"os"
	"flag"
	"dicorn"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: dicorn [run]\n")
	flag.PrintDefaults()
	os.Exit(1)
}

func main() {
	//os.Getenv("HOME") + "/.s3cfg")

	flag.Usage = usage
	flag.Parse()
	subcmd := flag.Arg(0)
	
	switch subcmd {
	case "run":
		dicorn.Run("0.0.0.0:9979")
	default:
		usage()
	}

}

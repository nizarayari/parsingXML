package main

import (
	"flag"
	"log"
)

var (
	listen  = flag.String("listen", ":9090", "Address to listen on")
	address = flag.String("address", "http://localhost:9090", "Address to listen on")
	mode    = flag.String("mode", "client", "Mode: server/client")
)

func main() {
	flag.Parse()
	switch *mode {
	case "client":
		client(*address)
	case "server":
		server(*listen)
	default:
		log.Fatalf("Unknow mode: %s", *mode)
	}
}

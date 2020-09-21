// Binary socks_proxy is a simple SOCKS5 proxy.
package main

import (
	"flag"
	"log"

	"github.com/armon/go-socks5"
)

var (
	flagListen = flag.String("listen", ":1080", "Address to listen on.")
)

func main() {
	flag.Parse()
	srv, err := socks5.New(&socks5.Config{})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Listening on %v", *flagListen)
	log.Fatal(srv.ListenAndServe("tcp", *flagListen))
}

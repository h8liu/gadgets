package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	addr := flag.String("addr", "localhost:8000", "listen address")
	dir := flag.String("dir", ".", "home directory")
	flag.Parse()
	log.Printf("listening on %s", *addr)
	http.Handle("/", http.FileServer(http.Dir(*dir)))
	log.Fatal(http.ListenAndServe(*addr, nil))
}

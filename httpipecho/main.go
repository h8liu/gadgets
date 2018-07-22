package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, req *http.Request) {
	addr, e := net.ResolveTCPAddr("tcp", req.RemoteAddr)
	if e != nil {
		fmt.Fprintf(w, "error: %s", addr)
	} else {
		fmt.Fprintf(w, addr.IP.String())
	}
}

func main() {
	port := flag.Int("port", 18888, "port to listen on")
	flag.Parse()

	fmt.Printf("httpipecho listen on port: %d\n", port)
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

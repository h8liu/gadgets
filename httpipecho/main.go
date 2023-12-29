package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
)

func handler(w http.ResponseWriter, req *http.Request) {
	addr, err := net.ResolveTCPAddr("tcp", req.RemoteAddr)
	if err != nil {
		errStr := fmt.Sprintf("invalid address: %s", err)
		http.Error(w, errStr, http.StatusBadRequest)
		return
	}

	fmt.Fprint(w, addr.IP.String())
}

func main() {
	port := flag.Int("port", 18888, "port to listen on")
	flag.Parse()

	fmt.Printf("httpipecho listen on port: %d\n", port)
	http.HandleFunc("/", handler)

	addr := fmt.Sprintf(":%d", port)

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}

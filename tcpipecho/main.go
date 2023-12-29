package main

import (
	"flag"
	"fmt"
	"log"
	"net"
)

func serve(c *net.TCPConn) {
	ret := c.RemoteAddr().(*net.TCPAddr).IP.String()
	if _, err := c.Write([]byte(ret)); err != nil {
		log.Print("write: ", err)
	}

	if err := c.Close(); err != nil {
		log.Print("close: ", err)
	}
}

func main() {
	port := flag.Int("port", 18889, "port to listen on")
	flag.Parse()

	iface := fmt.Sprintf(":%d", *port)
	addr, err := net.ResolveTCPAddr("tcp4", iface)
	if err != nil {
		log.Fatal("resolve TCP address: ", err)
	}

	lis, err := net.ListenTCP("tcp4", addr)
	if err != nil {
		log.Fatal("listen: ", err)
	}

	log.Printf("rawipecho listening on port: %d", port)

	for {
		con, err := lis.AcceptTCP()
		if err != nil {
			log.Fatal("accept: ", err)
		}

		go serve(con)
	}
}

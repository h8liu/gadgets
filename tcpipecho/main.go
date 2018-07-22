package main

import (
	"flag"
	"fmt"
	"log"
	"net"
)

func logError(e error) {
	if e != nil {
		log.Print(e)
	}
}

func noError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func serve(c *net.TCPConn) {
	ret := c.RemoteAddr().(*net.TCPAddr).IP.String()
	_, err := c.Write([]byte(ret))
	logError(err)
	err = c.Close()
	logError(err)
}

func main() {
	port := flag.Int("port", 18889, "port to listen on")
	flag.Parse()

	iface := fmt.Sprintf(":%d", *port)
	addr, err := net.ResolveTCPAddr("tcp4", iface)
	noError(err)

	lis, err := net.ListenTCP("tcp4", addr)
	noError(err)

	log.Printf("rawipecho listening on port: %d", port)

	for {
		con, err := lis.AcceptTCP()
		noError(err)

		go serve(con)
	}
}

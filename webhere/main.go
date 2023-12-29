package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
)

type flags struct {
	addr string
	dir  string
}

func run(ctx context.Context, f *flags, onReady func(addr string)) error {

	s := &http.Server{
		Handler: http.FileServer(http.Dir(f.dir)),
		Addr:    f.addr,
	}

	done := make(chan struct{})
	defer close(done)

	go func() {
		select {
		case <-ctx.Done():
			s.Close()
		case <-done:
		}
	}()

	lis, err := net.Listen("tcp", f.addr)
	if err != nil {
		return fmt.Errorf("listen: %w", err)
	}

	if onReady != nil {
		actualAddr := lis.Addr().String()
		log.Printf("listening on %s", actualAddr)
		onReady(actualAddr)
	}

	if err := s.Serve(lis); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}

func main() {
	f := new(flags)
	flag.StringVar(&f.addr, "addr", "localhost:8000", "listen address")
	flag.StringVar(&f.dir, "dir", ".", "home directory")
	flag.Parse()

	if err := run(context.Background(), f, nil); err != nil {
		log.Fatal(err)
	}
}

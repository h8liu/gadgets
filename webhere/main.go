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
	lis, err := net.Listen("tcp", f.addr)
	if err != nil {
		return fmt.Errorf("listen: %w", err)
	}

	if onReady != nil {
		actualAddr := lis.Addr().String()
		log.Printf("listening on %s", actualAddr)
		onReady(actualAddr)
	}

	s := &http.Server{Handler: http.FileServer(http.Dir(f.dir))}

	errChan := make(chan error)
	go func() { errChan <- s.Serve(lis) }()

	select {
	case <-ctx.Done():
		s.Close()
		<-errChan // Join the background server.

		// Only return error if it is cancelled context.
		if err := ctx.Err(); err != context.Canceled {
			return err
		}
	case err := <-errChan:
		if err != nil && err != http.ErrServerClosed {
			return err
		}
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

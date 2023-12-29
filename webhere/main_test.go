package main

import (
	"testing"

	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"time"
)

func TestRun(t *testing.T) {
	tmp := t.TempDir()

	index := []byte("<h1>hello</h1>")

	if err := os.WriteFile(
		filepath.Join(tmp, "index.html"), index, 0644,
	); err != nil {
		t.Fatal("write index page: ", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	f := &flags{
		addr: "127.0.0.1:0",
		dir:  tmp,
	}

	errChan := make(chan error, 1)
	addrChan := make(chan string)

	go func(ctx context.Context) {
		errChan <- run(ctx, f, func(addr string) { addrChan <- addr })
	}(ctx)

	var addr string
	select {
	case <-ctx.Done():
		t.Fatal(ctx.Err())
	case addr = <-addrChan:
	}

	c := &http.Client{}

	u := &url.URL{
		Scheme: "http",
		Host:   addr,
		Path:   "/",
	}

	resp, err := c.Get(u.String())
	if err != nil {
		t.Fatal("http get: ", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("read body: ", err)
	}

	if !bytes.Equal(body, index) {
		t.Errorf("got body %q, want %q", body, index)
	}

	cancel()

	if err := <-errChan; err != nil {
		t.Fatal("graceful close: ", err)
	}
}

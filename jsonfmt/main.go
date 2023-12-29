package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
)

func formatFile(in, out string) error {
	bs, err := os.ReadFile(in)
	if err != nil {
		return fmt.Errorf("read file: %w", err)
	}

	dest := new(bytes.Buffer)
	if err := json.Indent(dest, bs, "", "  "); err != nil {
		return fmt.Errorf("json encode: %w", err)
	}

	if err := os.WriteFile(out, dest.Bytes(), 0644); err != nil {
		return fmt.Errorf("create file: %w", err)
	}

	return nil
}

func main() {
	in := flag.String("in", "/dev/stdin", "input file")
	out := flag.String("out", "/dev/stdout", "output file")
	flag.Parse()
	args := flag.Args()

	if len(args) != 0 {
		for _, f := range args {
			if err := formatFile(f, f); err != nil {
				log.Fatal(err)
			}
		}
		return
	}

	if err := formatFile(*in, *out); err != nil {
		log.Fatal(err)
	}
}

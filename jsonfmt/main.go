package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"io"
	"log"
	"os"
)

func format(in, out string) error {
	bs, err := os.ReadFile(in)
	if err != nil {
		return err
	}

	dest := new(bytes.Buffer)
	if err := json.Indent(dest, bs, "", "  "); err != nil {
		return err
	}

	fout, err := os.Create(out)
	if err != nil {
		return err
	}
	defer fout.Close()

	if _, err := io.Copy(fout, dest); err != nil {
		return err
	}

	return fout.Close()
}

func main() {
	in := flag.String("in", "/dev/stdin", "input file")
	out := flag.String("out", "/dev/stdout", "output file")
	flag.Parse()
	args := flag.Args()

	if len(args) != 0 {
		for _, f := range args {
			if err := format(f, f); err != nil {
				log.Fatal(err)
			}
		}
		return
	}

	if err := format(*in, *out); err != nil {
		log.Fatal(err)
	}
}

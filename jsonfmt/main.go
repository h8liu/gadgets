package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func format(in, out string) error {
	bs, err := ioutil.ReadFile(in)
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

	if err := format(*in, *out); err != nil {
		log.Fatal(err)
	}
}

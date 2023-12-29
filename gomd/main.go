// Command gogfm converts an input markdown file into HTML.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"shanhu.io/g/markdown"
)

func main() {
	output := flag.String("out", "", "output file")
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		fmt.Fprintln(os.Stderr, "error: needs exactly one file")
		os.Exit(1)
	}

	bs, err := os.ReadFile(args[0])
	if err != nil {
		log.Fatal("read file: ", err)
	}

	result := markdown.ToHTML(bs)

	if *output != "" {
		if err := os.WriteFile(*output, result, 0644); err != nil {
			log.Fatal("write file: ", err)
		}
	} else {
		if _, err := os.Stdout.Write(result); err != nil {
			log.Fatal("write: ", err)
		}
	}
}

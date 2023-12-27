// Command gogfm converts an input markdown file into HTML.
package main

import (
	"flag"
	"fmt"
	"os"

	"shanhu.io/g/markdown"
)

func errExit(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func main() {
	output := flag.String("out", "", "output file")
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		fmt.Fprintln(os.Stderr, "error: needs exactly one file")
		os.Exit(1)
	}

	bs, err := os.ReadFile(args[0])
	errExit(err)

	result := markdown.ToHTML(bs)

	if *output != "" {
		errExit(os.WriteFile(*output, result, 0666))
	} else {
		_, err := os.Stdout.Write(result)
		errExit(err)
	}
}

// Command gogfm converts an input markdown file into HTML.
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	gfm "github.com/shurcooL/github_flavored_markdown"
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

	bs, err := ioutil.ReadFile(args[0])
	errExit(err)

	result := gfm.Markdown(bs)

	if *output != "" {
		errExit(ioutil.WriteFile(*output, result, 0666))
	} else {
		_, err := os.Stdout.Write(result)
		errExit(err)
	}
}

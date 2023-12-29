// Command gobytes reads a file and outputs its binary representation
// in golang. This could be useful for embedding binary resources
// into a Go language program.
package main

import (
	"fmt"
	"os"

	"flag"
)

func main() {
	var (
		varName  = flag.String("var", "v", "variable name")
		packName = flag.String("pack", "main", "package name")
		style    = flag.String("lang", "go", "language name")
	)
	flag.Parse()

	args := flag.Args()

	if len(args) != 1 {
		fmt.Fprintln(os.Stderr, "please specify exactly one file")
		os.Exit(1)
	}

	bytes, e := os.ReadFile(args[0])
	if e != nil {
		fmt.Fprintln(os.Stderr, e)
		os.Exit(1)
	}

	switch *style {
	case "go":
		fmt.Printf("package %s\n", *packName)
		fmt.Println()
		fmt.Printf("var %s = []byte{", *varName)

		for i, b := range bytes {
			if i%8 == 0 {
				fmt.Println()
				fmt.Print("\t")
			} else {
				fmt.Print(" ")
			}

			fmt.Printf("0x%02x,", b)
		}
		fmt.Println()
		fmt.Println("}")
	case "c":
		fmt.Printf("uint8_t %s[] = {", *varName)
		n := len(bytes)
		for i, b := range bytes {
			if i%8 == 0 {
				fmt.Println()
				fmt.Print("\t")
			} else {
				fmt.Print(" ")
			}
			fmt.Printf("0x%02x", b)
			if i != n-1 {
				fmt.Print(",")
			}
		}
		fmt.Println()
		fmt.Println("};")
	default:
		fmt.Fprintf(os.Stderr, "unknown language")
		os.Exit(1)
	}
}

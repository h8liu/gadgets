package main

import (
	"crypto/rand"
	"encoding/binary"
	"flag"
	"fmt"
	"log"
)

func makeRunes() []rune {
	var rs []rune
	for r := '0'; r <= '9'; r++ {
		rs = append(rs, r)
	}
	for r := 'a'; r <= 'z'; r++ {
		rs = append(rs, r)
	}
	for r := 'A'; r <= 'Z'; r++ {
		rs = append(rs, r)
	}
	rs = append(rs, '_', '-', '@', '!')

	return rs
}

func randUint32() uint32 {
	buf := make([]byte, 4)
	if _, err := rand.Reader.Read(buf); err != nil {
		log.Fatal(err)
	}
	return binary.LittleEndian.Uint32(buf)
}

func randPassword(n int) string {
	// TODO: while this is ok, would be nice if this function guarantees
	// that it always contains all types of characters.

	runes := makeRunes()
	nrune := uint32(len(runes))
	pwd := ""
	for i := 0; i < n; i++ {
		r := runes[randUint32()%nrune]
		pwd += string(r)
	}

	return pwd
}

func main() {
	n := flag.Int("n", 16, "password length")
	flag.Parse()

	fmt.Println(randPassword(*n))
}

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

func main() {
	length := flag.Int("n", 16, "password length")
	flag.Parse()

	runes := makeRunes()
	nrune := uint32(len(runes))
	pwd := ""
	for i := 0; i < *length; i++ {
		r := runes[randUint32()%nrune]
		pwd += string(r)
	}
	fmt.Println(pwd)
}

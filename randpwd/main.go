package main

import (
	"crypto/rand"
	"encoding/binary"
	"flag"
	"fmt"
	"log"
)

var runes = func() []rune {
	var ret []rune
	for r := '0'; r <= '9'; r++ {
		ret = append(ret, r)
	}
	for r := 'a'; r <= 'z'; r++ {
		ret = append(ret, r)
	}
	for r := 'A'; r <= 'Z'; r++ {
		ret = append(ret, r)
	}

	return ret
}()

func randUint32() uint32 {
	buf := make([]byte, 4)
	_, err := rand.Reader.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	return binary.LittleEndian.Uint32(buf)
}

func main() {
	length := flag.Int("n", 16, "password length")
	flag.Parse()

	nrune := uint32(len(runes))
	pwd := ""
	for i := 0; i < *length; i++ {
		r := runes[randUint32()%nrune]
		pwd += string(r)
	}
	fmt.Println(pwd)
}

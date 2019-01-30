package main

import (
	"bytes"
	"io"
	"log"
	"math/rand"
)

func main() {
	randbuffer := func(n int) *bytes.Buffer {
		rc := make([]byte, n)
		rand.Read(rc)
		return bytes.NewBuffer(rc)
	}
	// START OMIT
	src := randbuffer(32)
	buffSize := int(5)
	in := make([]byte, buffSize)

	for i := 0; ; i++ {
		nbytes, err := src.Read(in)
		log.Printf("%03d: nbytes = %d, err = %v", i, nbytes, err)

		if err == io.EOF {
			log.Printf("got EOF")
			break
		}
	}
	// END OMIT
}

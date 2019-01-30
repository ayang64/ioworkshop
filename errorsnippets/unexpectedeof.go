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
	src := randbuffer(1024) // allocates a bytes.Buffer of size N initalized with random bytes
	dst := make([]byte, 2048)

	if _, err := io.ReadAtLeast(src, dst, len(dst)); err != nil {
		log.Fatal(err)
	}
	// END OMIT
}

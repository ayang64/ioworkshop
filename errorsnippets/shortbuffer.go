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
	src := randbuffer(4096) // allocates a bytes.Buffer of size N initalized with random bytes
	dst := make([]byte, 2048)

	// attempt to read into a buffer that is too small to contain the "at least"
	// number of bytes.
	if _, err := io.ReadAtLeast(src, dst, len(src)); err != nil {
		log.Fatal(err)
	}
	// END OMIT
}

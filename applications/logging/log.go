package main

import (
	"flag"
	"io"
	"log"
	"math/rand"
	"os"
	"time"
)

type ErrWriter struct {
	w   io.Writer
	err error
}

func (ew *ErrWriter) Write(p []byte) (int, error) {
	if ew.err != nil {
		return 0, ew.err
	}
	if rand.Intn(1000) == 0 {
		ew.err = io.EOF
	}
	return ew.w.Write(p)
}

func main() {
	logFile := flag.String("l", "", "path to log file.")
	logStderr := flag.Bool("v", false, "log to os.Stderr")
	logQuiet := flag.Bool("q", false, "be quiet -- log nothing regarless of other parameters")
	flag.Parse()

	rand.Seed(time.Now().UnixNano())

	writers := func() (w []io.Writer) {
		if *logQuiet {
			return
		}
		if *logFile != "" {
			fh, err := os.Create(*logFile)
			if err != nil {
				log.Fatal(err)
			}
			w = append(w, &ErrWriter{w: fh})
		}
		if *logStderr {
			w = append(w, os.Stderr)
		}
		return
	}

	// it is cool that io.MultiWriter() effectively discards data if it has an
	// empty writer list.
	l := log.New(io.MultiWriter(writers()...), "DEBUG: ", log.LstdFlags|log.Ldate|log.Lshortfile)

	rnd := make([]byte, 32)
	for {
		// generate some random data.
		rand.Read(rnd)
		l.Printf("%x", rnd)
	}
}

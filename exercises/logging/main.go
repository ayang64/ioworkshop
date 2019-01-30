package main

import (
	"flag"
	"io"
	"log"
	"os"
)

func main() {
	quiet := flag.Bool("q", false, "the program should not generate any output.")
	logfile := flag.String("l", "", "path to log file.")
	flag.Parse()

	writers := func() ([]io.Writer, error) {
		if *quiet {
			return nil, nil
		}
		rc := []io.Writer{}
		if *logfile != "" {
			f, err := os.Create(*logfile)
			if err != nil {
				return nil, err
			}
			rc = append(rc, f)
		}
		return append(rc, os.Stderr), nil
	}
	w, err := writers()

	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(io.MultiWriter(w...))

	log.Printf("hello, world!")
}

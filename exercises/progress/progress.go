package main

import (
	"bytes"
	"io"
	"log"
	"math/rand"
	"time"

	"github.com/ayang64/ioworkshop/exercises/slowwriter"
)

type Spinner interface {
	Spin(int, int)
}

type progress struct {
	max     int
	cur     int
	spinner Spinner
}

type DefaultSpinner struct{}

func (DefaultSpinner) Spin(cur int, max int) {
	log.Printf("%d of %d bytes transferred (%.03f%%)", cur, max, (float64(cur)/float64(max))*100.0)
}

func (p *progress) Write(b []byte) (int, error) {
	p.cur += len(b)
	p.spinner.Spin(p.cur, p.max)
	return len(b), nil
}

func NewWriter(w io.Writer, max int) io.Writer {
	p := progress{
		max:     max,
		cur:     0,
		spinner: DefaultSpinner{},
	}
	return io.MultiWriter(w, &p)
}

func NewReader(r io.Reader, max int) io.Reader {
	p := progress{
		max:     max,
		cur:     0,
		spinner: DefaultSpinner{},
	}
	return io.TeeReader(r, &p)
}

func main() {
	size := 1 << 26
	rand.Seed(time.Now().UnixNano())
	randbytes := make([]byte, size)
	rand.Read(randbytes)
	src := bytes.NewBuffer(randbytes)

	dst := bytes.NewBuffer(make([]byte, size))

	io.Copy(slowwriter.New(dst, 14400), NewReader(src, size))
}

package slowwriter

import (
	"io"
	"time"
)

type slowWriter struct {
	delay time.Duration
	w     io.Writer
}

func (s slowWriter) Write(b []byte) (int, error) {
	time.Sleep(s.delay * time.Duration(len(b)))
	return s.w.Write(b)
}

func New(w io.Writer, bps int) io.Writer {
	return &slowWriter{
		delay: time.Duration(1e+09 / float64(bps*8)),
		w:     w,
	}
}

package repeater_test

import (
	"bytes"
	"github.com/ayang64/ioworkshop/exercises/repeater"
	"io"
	"testing"
)

func TestRepeaterRead(t *testing.T) {
	bb := &bytes.Buffer{}

	rptr, err := repeater.New([]byte("ABCDEFG"))

	if err != nil {
		t.Fatal(err)
	}

	r := io.LimitReader(rptr, 1025)

	io.Copy(bb, r)

	t.Log(bb.String())
}

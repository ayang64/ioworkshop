package essthree_test

import (
	"fmt"
	"io"
	"os"
	"testing"
	"time"

	"github.com/ayang64/ioworkshop/applications/essthree"
)

func TestCopyTraditional(t *testing.T) {
	if err := essthree.CopyTraditional("/etc/passwd", "passwd2"); err != nil {
		t.Fatal(err)
	}
}

func TestMultiCreate(t *testing.T) {
	dests := map[string]func(string) (io.Writer, error){
		"s3":   essthree.Create,
		"file": func(s string) (io.Writer, error) { return os.Create(s) },
	}

	for k, open := range dests {
		func() {
			t.Logf("writing to %s", k)
			w, err := open("test-fixtures/test-file")
			if err != nil {
				t.Error(err)
			}
			i, err := os.Open("/etc/rc.conf")
			io.Copy(w, i)
			if c, ok := w.(io.Closer); ok {
				c.Close()
			}
		}()
	}
}

func TestCreate(t *testing.T) {
	w, err := essthree.Create("rc.conf")

	if err != nil {
		t.Fatal(err)
	}

	inf, _ := os.Open("/etc/rc.conf")

	if _, err := io.Copy(w, inf); err != nil {
		t.Fatal(err)
	}

	for i := 0; i < 1000; i++ {
		fmt.Fprintf(w, "TEST: %d", i)
	}

	if c, ok := w.(io.Closer); ok {
		c.Close()
	}
	time.Sleep(2 * time.Second)
}

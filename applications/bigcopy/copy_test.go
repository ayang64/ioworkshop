package copy

import (
	"fmt"
	"os"
	"path"
	"testing"
)

func setup(file string, size int64) (func() error, error) {
	dir := "test-fixtures"

	os.RemoveAll(dir)

	p := path.Join(dir, file)
	if err := os.Mkdir(dir, 0755); err != nil {
		return nil, err
	}
	fh, err := os.Create(path.Join(dir, file))

	if err != nil {
		return nil, err
	}
	fh.Close()

	os.Truncate(p, size)

	cleanup := func() error {
		os.RemoveAll(dir)
		return nil
	}
	return cleanup, nil
}

func BenchmarkBigCopy(b *testing.B) {
	benches := []struct {
		Name   string
		MinExp uint // minimum buffer size epxonent
		MaxExp uint // maximum buffer size exponent
		F      func(string, string, []byte) (int64, error)
	}{
		{Name: "CopyInternalBuffer", MinExp: 1, MaxExp: 2, F: CopyUnBuffered},
		{Name: "CopyExternalBuffer", MinExp: 4, MaxExp: 24, F: CopyBuffered},
	}

	cleanup, err := setup("bigfile", 1<<24)
	if err != nil {
		b.Fatal(err)
	}
	defer cleanup()
	for _, bench := range benches {
		for i := bench.MinExp; i < bench.MaxExp; i++ {
			b.Run(bench.Name, func(b *testing.B) {
				bufsize := 1 << i
				buf := make([]byte, bufsize)
				b.Run(fmt.Sprintf("%08d", bufsize), func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						bench.F("test-fixtures/bigfile", "test-fixtures/copiedfile", buf)
					}
				})
			})
		}
	}
}

package copy

import (
	"io"
	"os"
)

func CopyUnBuffered(src string, dst string, buf []byte) (int64, error) {
	in, err := os.Open(src)

	if err != nil {
		return 0, err
	}

	out, err := os.Create(dst)

	if err != nil {
		return 0, err
	}

	return io.Copy(out, in)
}

func CopyBuffered(src string, dst string, buf []byte) (int64, error) {
	in, err := os.Open(src)

	if err != nil {
		return 0, err
	}

	out, err := os.Create(dst)

	if err != nil {
		return 0, err
	}

	return io.CopyBuffer(out, in, buf)
}

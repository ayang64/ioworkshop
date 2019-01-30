package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	openall := func(paths ...string) ([]io.Reader, []error) {
		if len(paths) == 0 {
			return []io.Reader{os.Stdin}, []error(nil)
		}
		rc := []io.Reader{}
		errors := []error{}
		for i := range paths {
			fh, err := os.Open(paths[i])
			if err != nil {
				errors = append(errors, err)
			}
			rc = append(rc, fh)
		}
		return rc, errors
	}

	closeall := func(readers ...io.Reader) error {
		for i := range readers {
			if closer, ok := readers[i].(io.Closer); ok {
				closer.Close()
			}
		}
		return nil
	}

	readers, openerrors := openall(os.Args[1:]...)
	defer closeall(readers...)

	io.Copy(os.Stdout, io.MultiReader(readers...))

	for i := range openerrors {
		fmt.Fprintln(os.Stderr, openerrors[i])
	}
}

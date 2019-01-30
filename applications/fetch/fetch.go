package main

import (
	"crypto/sha256"
	"fmt"
	"hash"
	"io"
	"log"
	"net/http"
	"os"
	"path"
)

func fetchandhash(u string, destpath string) (hash.Hash, error) {
	resp, err := http.Get(u)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("%q returned status %s", u, resp.Status)
	}
	defer resp.Body.Close()

	w, err := os.Create(destpath)
	if err != nil {
		return nil, err
	}
	defer w.Close()

	return copysum(w, resp.Body)
}

func copysum(w io.Writer, r io.Reader) (hash.Hash, error) {
	hash := sha256.New()
	io.Copy(io.MultiWriter(o, hash), resp.Body)
	return hash, nil
}

func main() {
	for _, url := range os.Args[1:] {
		dest := path.Base(url)
		hash, err := fetchandhash(url, dest)
		if err != nil {
			log.Print(err)
		}
		fmt.Printf("%x - %s\n", hash.Sum(nil), dest)
	}
}

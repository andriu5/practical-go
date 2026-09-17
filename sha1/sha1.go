package main

import (
	"compress/gzip"
	"crypto/sha1"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fmt.Println(SHA1Sig("http.log.gz"))
	fmt.Println(SHA1Sig("sha1.go"))
}

// SHA1Sig returns SHA1 signature of uncompressed file.
// Exercise: Decompress only if file name ends with ".gz"
// cat http.log.gz | gunzip | sha1sum
func SHA1Sig(fileName string) (string, error) {

	// cat http.log.gz
	file, err := os.Open(fileName)
	if err != nil {
		return "", fmt.Errorf("%q - open: %w", fileName, err)
	}

	defer file.Close()

	var r io.Reader = file

	// | gunzip
	if strings.HasSuffix(fileName, ".gz") {
		// BUG: Creates new "r" that is only in "if" scope.
		// this is called shadowing
		// r, err := gzip.NewReader(file)
		gz, err := gzip.NewReader(file)
		if err != nil {
			return "", fmt.Errorf("%q - gzip: %w", fileName, err)
		}
		defer gz.Close()
		r = gz
	}

	// | sha1sum
	w := sha1.New()
	if _, err := io.Copy(w, r); err != nil {
		return "", fmt.Errorf("%q - copy: %w", fileName, err)
	}

	sig := w.Sum(nil)
	return fmt.Sprintf("%x", sig), nil
}

/*
Go
type Reader interface {
	Read(p []byte) (n int, err error)
}

Python
type Reader interface {
	Read(n int) ([]byte, err error)
}
*/

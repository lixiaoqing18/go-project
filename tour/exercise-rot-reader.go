package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (reader rot13Reader) Read(bytes []byte) (int, error) {
	n, e := reader.r.Read(bytes)
	for i := 0; i < len(bytes); i++ {
		if bytes[i] >= 'A' && bytes[i] < 'N' || bytes[i] >= 'a' && bytes[i] < 'n' {
			bytes[i] = bytes[i] + 13
		} else if bytes[i] > 'M' && bytes[i] <= 'Z' || bytes[i] > 'm' && bytes[i] <= 'z' {
			bytes[i] = bytes[i] - 13
		}
	}
	return n, e
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

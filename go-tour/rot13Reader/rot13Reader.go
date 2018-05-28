package main

import (
	"io"
	"os"
	"strings"
	//"fmt"
)

type rot13Reader struct {
	r io.Reader
}

func (reader rot13Reader) Read(result []byte) (int, error) {
	// Read from the unciphered reader:
	nBytes, error := reader.r.Read(result)
	
	for i := 0; i < nBytes; i++ {
		if (result[i] >= 'A' && result[i] <= 'Z') {
			result[i] = 'A' + ((result[i] - 'A' + 13) % 26)
		}
		
		if (result[i] >= 'a' && result[i] <= 'z') {
			result[i] = 'a' + ((result[i] - 'a' + 13) % 26)
		}
	}
	
	return nBytes, error
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

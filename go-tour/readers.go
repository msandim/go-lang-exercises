package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (reader MyReader) Read(result []byte) (int, error) {
	for i := range result {
		result[i] = 'A'
	}
	
	return len(result), nil
}

func main() {
	reader.Validate(MyReader{})
}

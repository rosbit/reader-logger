package logr

import (
	"encoding/json"
	"testing"
	"bytes"
	"os"
	"io"
	"fmt"
)

func TestReaderLogger(t *testing.T) {
	reader := bytes.NewBufferString(`{"name":"rosbit", "age": 10}`)

	var j struct {
		Name string `json:"name"`
		Age int `json:"age"`
	}

	if err := parseJSON(reader, &j); err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Printf("result: %v\n", j)
}

func parseJSON(reader io.Reader, j interface{}) error {
	r, deferFunc := ReaderLogger(reader, os.Stderr, "data in reader")
	defer deferFunc()

	return json.NewDecoder(r).Decode(j)
}

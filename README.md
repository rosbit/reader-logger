# reader-logger, a utility to output the reader content and process it simultaneously.

## usage

```go
import (
	logr "github.com/rosbit/reader-logger"
	"encoding/json"
	"bytes"
	"os"
	"io"
	"fmt"
)

func main() {
	// a reader for testing
	reader := bytes.NewBufferString(`{"name":"rosbit", "age": 10}`)

	// j to store the result
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

// print the content in reader and parse it.
func parseJSON(reader io.Reader, j interface{}) error {
	r, deferFunc := logr.ReaderLogger(reader, os.Stderr, "data in reader") // create a new reader
	defer deferFunc()

	return json.NewDecoder(r).Decode(j) // parse JSON in the new reader
}
```

## result

run `go test`, the result is as following:

```
--- data in reader begin ---
{"name":"rosbit", "age": 10}
--- data in reader end ---
result: {rosbit 10}
PASS
ok  	github.com/rosbit/reader-logger	0.685s
```

you can output the content of reader in any io.Writer other than os.Stderr as in the sample, of course.

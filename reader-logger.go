package logr

import (
	"fmt"
	"io"
)

func ReaderLogger(reader io.Reader, logger io.Writer, prompt string) (newReader io.Reader, deferFunc func()) {
	if logger == nil {
		return reader, func(){}
	}

	fmt.Fprintf(logger, "--- %s begin ---\n", prompt)
	return io.TeeReader(reader, logger), func() {
		fmt.Fprintf(logger, "\n--- %s end ---\n", prompt)
	}
}

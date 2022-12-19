package logr

import (
	"fmt"
	"io"
)

func ReaderLogger(reader io.Reader, logger io.Writer, prompt string) (newReader io.Reader, deferFunc func()) {
	if logger == nil {
		return reader, func(){}
	}

	newReader = io.TeeReader(reader, logger)
	if len(prompt) > 0 {
		fmt.Fprintf(logger, "--- %s begin ---\n", prompt)
		deferFunc = func() {
			fmt.Fprintf(logger, "\n--- %s end ---\n", prompt)
		}
	} else {
		deferFunc = func(){}
	}

	return
}

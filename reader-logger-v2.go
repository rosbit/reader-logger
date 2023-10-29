package logr

import (
	"fmt"
	"io"
)

func ReaderLogger2(reader io.Reader, logger io.Writer, prompt ...string) (newReader io.Reader, beginFunc, endFunc func()) {
	if logger == nil {
		return reader, func(){}, func(){}
	}

	newReader = io.TeeReader(reader, logger)
	if len(prompt) > 0 && len(prompt[0]) > 0 {
		beginFunc = func() {
			fmt.Fprintf(logger, "--- %s begin ---\n", prompt[0])
		}
		endFunc = func() {
			fmt.Fprintf(logger, "\n--- %s end ---\n", prompt[0])
		}
	} else {
		beginFunc = func(){}
		endFunc = func(){}
	}

	return
}

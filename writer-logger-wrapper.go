package logr

import (
	"net/http"
	"io"
	"fmt"
)

func CreateResponseWriter(w http.ResponseWriter, logger io.Writer, prompt ...string) (newW http.ResponseWriter, endFunc func()) {
	if logger == nil {
		return w, func(){}
	}

	newWriter := io.MultiWriter(w, logger)
	var beginFunc func()
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

	newW = wrapRespWriter(newWriter, w, beginFunc)
	return
}

type respWriter struct {
	w http.ResponseWriter
	newWriter io.Writer
	beginFunc func()
	firstWriting bool
}
func (w *respWriter) Header() http.Header {
	return w.w.Header()
}
func (w *respWriter) Write(c []byte) (int, error) {
	if w.firstWriting {
		w.firstWriting = false
		w.beginFunc()
	}
	return w.newWriter.Write(c)
}
func (w *respWriter) WriteHeader(statusCode int) {
	w.w.WriteHeader(statusCode)
}
func wrapRespWriter(newWriter io.Writer, w http.ResponseWriter, beginFunc func()) *respWriter {
	return &respWriter{
		w: w,
		newWriter: newWriter,
		beginFunc: beginFunc,
		firstWriting: true,
	}
}

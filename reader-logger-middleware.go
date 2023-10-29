package logr

import (
	"net/http"
	"io"
	"os"
)

type HttpHandlerFunc = func(http.ResponseWriter, *http.Request, http.HandlerFunc)

func CreateBodyDumpingHandlerFunc(dumper io.Writer, prompts ...string) HttpHandlerFunc {
	var prompt string
	if len(prompts) > 0 {
		prompt = prompts[0]
	}
	if dumper != nil {
		if f, ok := dumper.(*os.File); ok {
			if f == os.Stderr || f == os.Stdout {
				if len(prompt) == 0 {
					prompt = "dumping body"
				}
			}
		}
	}

	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		if r.Body == nil {
			next(w, r)
			return
		}
		nr, beginFunc, endFunc := ReaderLogger2(r.Body, dumper, prompt)
		r.Body = wrapNopCloser(nr, beginFunc, endFunc)
		next(w, r)
	}
}

type nopCloser struct {
	io.Reader
	beginFunc, endFunc func()
	firstReading bool
	bfCalled, efCalled bool
}
func (rc *nopCloser) Read(p []byte) (n int, err error) {
	if !rc.bfCalled {
		rc.bfCalled = true
		rc.beginFunc()
	}
	if n, err = rc.Reader.Read(p); err != nil && !rc.efCalled {
		if err == io.EOF {
			rc.efCalled = true
			rc.endFunc()
		}
	}
	return
}
func (rc *nopCloser) Close() error {
	if !rc.efCalled {
		rc.efCalled = true
		rc.endFunc()
	}
	if c, ok := rc.Reader.(io.ReadCloser); ok {
		return c.Close()
	}
	return nil
}
func wrapNopCloser(r io.Reader, beginFunc, endFunc func()) *nopCloser {
	return &nopCloser{
		Reader: r,
		beginFunc: beginFunc,
		endFunc: endFunc,
	}
}

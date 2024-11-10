package Http

import (
	"fmt"
	"syscall"
)

type Header map[string][]string

func (h Header) Set(key, value string) {
	h[key] = []string{value}
}

type Response struct {
	fd          int
	header      Header
	status      int
	headersSent bool
}

func (w *Response) Header() Header {
	return w.header
}

func (w *Response) WriteHeader(statusCode int) {
	if !w.headersSent {
		w.status = statusCode
		w.Write(nil)
	}
}

func (w *Response) Write(data []byte) (int, error) {
	if w.status == 0 {
		w.status = StatusOK
	}
	if !w.headersSent {
		statusLine := fmt.Sprintf("HTTP/1.1 %d %s\r\n", w.status, GetStatusText(w.status))
		_, err := syscall.Write(w.fd, []byte(statusLine))
		if err != nil {
			return 0, err
		}

		for key, values := range w.header {
			for _, value := range values {
				_, err := syscall.Write(w.fd, []byte(fmt.Sprintf("%s: %s\r\n", key, value)))
				if err != nil {
					return 0, err
				}
			}
		}

		_, err = syscall.Write(w.fd, []byte("\r\n"))
		if err != nil {
			return 0, err
		}
		w.headersSent = true
	}

	return syscall.Write(w.fd, data)
}

func NewResponseWriter(fd int) *Response {
	return &Response{
		fd:          fd,
		header:      make(Header),
		status:      0,
		headersSent: false,
	}
}

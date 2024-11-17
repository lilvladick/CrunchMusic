package Http

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"net/url"
	"strings"
)

type Request struct {
	Method string
	URL    *url.URL
	Proto  string
	Header map[string][]string
	Body   string
	ctx    context.Context
}

func ReadRequest(reader *bufio.Reader) (*Request, error) {
	firstLine, err := reader.ReadString('\n')
	if err != nil {
		return nil, fmt.Errorf("не удалось прочитать первую строку: %v", err)
	}

	firstLine = strings.TrimSpace(firstLine)
	parts := strings.Fields(firstLine)
	if len(parts) < 3 {
		return nil, fmt.Errorf("неверный формат запроса")
	}

	method := parts[0]
	urlString := parts[1]
	proto := parts[2]

	parsedURL, err := url.Parse(urlString)
	if err != nil {
		return nil, fmt.Errorf("ошибка парсинга URL: %v", err)
	}

	req := &Request{
		Method: method,
		URL:    parsedURL,
		Proto:  proto,
		Header: make(map[string][]string),
	}

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return nil, fmt.Errorf("ошибка при чтении заголовков: %v", err)
		}
		line = strings.TrimSpace(line)
		if line == "" {
			break
		}
		headerParts := strings.SplitN(line, ": ", 2)
		if len(headerParts) == 2 {
			key := headerParts[0]
			value := headerParts[1]
			req.Header[key] = append(req.Header[key], value)
		}
	}

	var body bytes.Buffer
	if _, err := body.ReadFrom(reader); err != nil && err != io.EOF {
		return nil, fmt.Errorf("ошибка при чтении тела: %w", err)
	}
	req.Body = body.String()

	return req, nil
}

func (r *Request) Context() context.Context {
	if r.ctx != nil {
		return r.ctx
	}
	return context.Background()
}

func (r *Request) WithContext(ctx context.Context) *Request {
	if ctx == nil {
		panic("nil context")
	}
	r2 := new(Request)
	*r2 = *r
	r2.ctx = ctx
	return r2
}

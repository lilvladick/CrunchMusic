package router

import (
	"CrunchServer/handlers"
	Http "CrunchServer/http"
	"bufio"
	"bytes"
	"context"
	"fmt"
	"log"
	"regexp"
	"strings"
	"syscall"
)

type HandlerFunc func(Http.Response, *Http.Request)

var routes = []route{
	NewRoute("GET", "/authors", handlers.AllAuthors),
	NewRoute("POST", "/authors", handlers.AddAuthor),
	NewRoute("GET", "/authorbyemail", handlers.GetAuthorByEmail),
	NewRoute("GET", "/authorbyid", handlers.GetAuthorById),
	NewRoute("GET", "/authorsbyname", handlers.GetAuthorByName),
	NewRoute("GET", "/category", handlers.AllCategories),
	NewRoute("POST", "/category", handlers.AddCategory),
	NewRoute("GET", "/categorybyid", handlers.GetCategoryById),
	NewRoute("GET", "/categorybyname", handlers.GetCategoryByName),
	NewRoute("GET", "/news", handlers.AllNews),
	NewRoute("POST", "/news", handlers.WriteNews),
	NewRoute("GET", "/newsbyid", handlers.GetNewsById),
	NewRoute("GET", "/newsbyauthorid", handlers.GetNewsByAuthorId),
	NewRoute("GET", "/newsbycategoryid", handlers.GetNewsByCategoryId),
	NewRoute("GET", "/newsbytitle", handlers.GetNewsByTitle),
	NewRoute("GET", "/comments", handlers.AllComments),
	NewRoute("POST", "/comments", handlers.WriteComment),
	NewRoute("GET", "/commentsbyauthorid", handlers.GetCommentByAuthorId),
	NewRoute("GET", "/commentsbyid", handlers.GetCommentById),
	NewRoute("GET", "/commentsbynewsid", handlers.GetCommentByNewsId),
}

func NewRoute(method, pattern string, handler HandlerFunc) route {
	return route{method, regexp.MustCompile("^" + pattern + "$"), handler}
}

type route struct {
	method  string
	regex   *regexp.Regexp
	handler HandlerFunc
}

func Serve(w Http.Response, r *Http.Request) {
	var allow []string
	for _, route := range routes {
		matches := route.regex.FindStringSubmatch(r.URL.Path)
		if len(matches) > 0 {
			if r.Method != route.method {
				allow = append(allow, route.method)
				continue
			}
			ctx := context.WithValue(r.Context(), ctxKey{}, matches[1:])
			route.handler(w, r.WithContext(ctx))
			return
		}
	}
	if len(allow) > 0 {
		w.Header().Set("Allow", strings.Join(allow, ", "))
		w.WriteHeader(Http.StatusMethodNotAllowed) // 405 Method Not Allowed
		fmt.Println(Http.GetStatusText(Http.StatusMethodNotAllowed))
		return
	}
	w.WriteHeader(Http.StatusNotFound) // 404 Not Found
	fmt.Println(Http.GetStatusText(Http.StatusNotFound))
}

func ListenAndServe() error {
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0) // AF_INET = IPv4, SOCK_STREAM = TCP
	if err != nil {
		log.Print("socket creating errorrror: ", err)
		return err
	}
	defer syscall.Close(fd)

	addr := syscall.SockaddrInet4{Port: 8080}
	copy(addr.Addr[:], []byte{0, 0, 0, 0})

	err = syscall.Bind(fd, &addr)
	if err != nil {
		log.Print("bind error: ", err)
		return err
	}

	err = syscall.Listen(fd, 10)
	if err != nil {
		log.Print("listen error: ", err)
		return err
	}
	log.Print("server started")

	for {
		connFd, _, err := syscall.Accept(fd) // AF_INET = IPv4, SOCK_STREAM = TCP
		if err != nil {
			log.Print("accept error: ", err)
			return err
		}

		go HandleFd(connFd)
	}

}

func HandleFd(fd int) {
	defer syscall.Close(fd)
	buf := make([]byte, 1024)
	n, err := syscall.Read(fd, buf)
	if err != nil {
		log.Print("read error: ", err)
		return
	}
	reader := bytes.NewBuffer(buf[:n])

	req, err := Http.ReadRequest(bufio.NewReader(reader))
	if err != nil {
		log.Print("error: ", err)
		return
	}

	res := Http.NewResponseWriter(fd)
	Serve(*res, req)
}

type ctxKey struct{}

// func GetField(r *http.Request, index int) string {
// 	fields := r.Context().Value(ctxKey{}).([]string)
// 	return fields[index]
// }

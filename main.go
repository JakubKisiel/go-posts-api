package main

import (
	"io"
	"net/http"
)

type HttpHandler struct{}

func (h HttpHandler) ServeHTTP(res http.ResponseWriter, r *http.Request) {
	io.WriteString(res, "Hello, world!\n")
}

func main() {
	http.ListenAndServe(":9000", HttpHandler{})
}

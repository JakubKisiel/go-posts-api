package main

import (
	"io"
	"net/http"
	"sync"
)

type HttpHandler struct{}

func (h HttpHandler) ServeHTTP(res http.ResponseWriter, r *http.Request) {
	io.WriteString(res, "Hello, world!\n")
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		http.ListenAndServe(":9000", HttpHandler{})
	}()
}

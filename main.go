package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

type HttpHandler struct{}

func (h HttpHandler) ServeHTTP(res http.ResponseWriter, r *http.Request) {
	io.WriteString(res, "Hello, world!\n")
}

func ServeRoot(res http.ResponseWriter, r *http.Request) {
	io.WriteString(res, "Hello, "+r.Host)
}
func ServeTime(res http.ResponseWriter, r *http.Request) {
	io.WriteString(res, "Hello "+r.Host+" It's "+time.Now().Format("2006-01-02T15:04"))
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Starting Http server")
		http.HandleFunc("/", ServeRoot)
		http.HandleFunc("/time/", ServeTime)
		http.ListenAndServe(":9000", nil)
	}()
	wg.Wait()
	fmt.Println("Server is closed")
}

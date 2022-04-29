package main

import "net/http"

type HttpHandler struct{}

func (h HttpHandler) ServeHTTP(res http.ResponseWriter, r *http.Request) {
	data := []byte("Hello, world\n")
	res.Write(data)
}

func main() {
	http.ListenAndServe(":9000", HttpHandler{})
}

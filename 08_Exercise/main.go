package main

import (
	"io"
	"net/http"
)

func handler(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "URL: "+req.URL.Path)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

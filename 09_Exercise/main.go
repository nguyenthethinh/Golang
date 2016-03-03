package main

import (
	"net/http"
	"fmt"
	"strings"
	"io"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	name := strings.Split(r.URL.Path, "/")
	fmt.Println(name)
	io.WriteString(w, name[1])

}


func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

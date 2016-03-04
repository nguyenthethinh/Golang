package main

import(
	"net/http"
	"fmt"
	"io"
)

func main(){
	http.HandleFunc("/", func(res http.ResponseWriter, req * http.Request) {
		fmt.Println(req.URL.Path)
		fmt.Println(req.URL.RequestURI())
		io.WriteString(res, "Name: " + req.FormValue("n"))
	})
	http.ListenAndServe(":8080", nil)
}

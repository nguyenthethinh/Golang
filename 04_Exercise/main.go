package main

import (
	"net/http"
	"text/template"
)


func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("template/index.gohtml")
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("style/css"))))
	http.Handle("/pic/", http.StripPrefix("/pic", http.FileServer(http.Dir("style/pic"))))
	http.ListenAndServe(":8080", nil)
}

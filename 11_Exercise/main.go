package main

import (
	"html/template"
	"net/http"
	"fmt"
)

type Person struct {
	FirstName string
	LastName  string
}

func renderTemplate(res http.ResponseWriter, p *Person) {
	t, err := template.ParseFiles("name.html")
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(res, p)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

func displayHandler(res http.ResponseWriter, req *http.Request) {
	//title := r.URL.Path[len("/save/"):]
	fn := req.FormValue("firstName")
	ln := req.FormValue("lastName")
	p := &Person{FirstName: fn, LastName: ln}
	if req.Method == "POST" {
		fmt.Println("Form is posted....:" + p.FirstName + " " + p.LastName)
		renderTemplate(res, p)
	}
}

func main() {
	http.HandleFunc("/", displayHandler)
	http.ListenAndServe(":8080", nil)
}

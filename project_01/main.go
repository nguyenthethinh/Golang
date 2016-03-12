package main

import (
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request){
	tpl := template.Must(template.New("main").ParseGlob("*.html"))
	tplVars := map[string]string {
		"Title": "Template Title",
		"Content": "Content of template",
	}
	err := tpl.ExecuteTemplate(w, "index.html", tplVars)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)

}

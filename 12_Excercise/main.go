package main

import (
	"html/template"
	"net/http"
	"io"
	"os"
	"io/ioutil"
	"fmt"
)

type FileContent struct{
	Content string
}

func renderTemplate(res http.ResponseWriter, cont *FileContent) {
	t, err := template.ParseFiles("uploadFile.html")
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(res, cont)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
}

func loadForm(res http.ResponseWriter, req *http.Request){
	fmt.Println("Loading form.....")
	renderTemplate(res, &FileContent{""})
}

func uploadHandler(res http.ResponseWriter, req *http.Request) {
	// the FormFile function takes in the POST input id file
	if req.Method == "POST" {
		fmt.Println("Upload handling....")
		fmt.Println(req.URL.Path)
		file, _, err := req.FormFile("fileName")

		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

		defer file.Close()

		out, err := os.Create("uploadedfile.txt")
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

		defer out.Close()

		// write the content from POST to the file
		_, err = io.Copy(out, file)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
		}

		conArray, err := ioutil.ReadFile("uploadedfile.txt")
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
		}
		fc := &FileContent{string(conArray)}
		renderTemplate(res, fc)
	} else{
		fmt.Println("Why don't you post?")
	}
}

func main() {
	http.HandleFunc("/", loadForm)
	http.HandleFunc("/upload", uploadHandler)
	http.ListenAndServe(":8080", nil)
}

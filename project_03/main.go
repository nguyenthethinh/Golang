package main

import (
	"net/http"
	"html/template"
	"strconv"
)

type Person struct {
	Name string
	Age  int
}


func displayHandler(res http.ResponseWriter, req *http.Request) {
	tpl := template.Must(template.New("main").ParseGlob("*.html"))
	tplVars := map[string]string {
		"Cookie": "",
	}
	err := tpl.ExecuteTemplate(res, "index.html", tplVars)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

func cookieHandler(res http.ResponseWriter, req *http.Request){
	// Generating a new ID
	name := req.FormValue("name")
	age, _ := strconv.Atoi(req.FormValue("age"))
	p := Person{Name: name, Age: age}
	if req.Method == "POST" {
		cookie, err := req.Cookie("my-cookie")
		if err == http.ErrNoCookie{
			cookie = &http.Cookie{
				Name: "my-cookie",
				Value: p.Name + strconv.Itoa(p.Age),
				HttpOnly:true,
				//Secure: true,
			}
		}
		// Setting the cookie on the response back to the client
		http.SetCookie(res, cookie)
		tpl := template.Must(template.New("main").ParseGlob("*.html"))
		tplVars := map[string]string {
			"Cookie": cookie.Value,
		}
		err = tpl.ExecuteTemplate(res, "index.html", tplVars)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
		}
	}
}

func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", displayHandler)
	http.HandleFunc("/setcookie", cookieHandler)
	http.ListenAndServe(":8080", nil)
}

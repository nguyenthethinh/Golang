package main

import (
	"net/http"
	"github.com/nu7hatch/gouuid"

"html/template"
)


func cookieHandler(res http.ResponseWriter, req *http.Request){
	// Generating a new ID
	cookie, err := req.Cookie("session-id")
	if err == http.ErrNoCookie{
		id, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name: "session-id",
			Value: id.String(),
			HttpOnly:true,
			//Secure: true,
		}
	}
	// Setting the cookie on the response back to the client
	http.SetCookie(res, cookie)
	tpl := template.Must(template.New("main").ParseGlob("*.html"))
	tplVars := map[string]string {
		"Title": "SessionID",
		"Content": cookie.Value,
	}
	err = tpl.ExecuteTemplate(res, "index.html", tplVars)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", cookieHandler)
	http.ListenAndServe(":8080", nil)
}

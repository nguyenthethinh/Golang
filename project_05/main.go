package main

import (
	"net/http"
	"html/template"
	"strconv"
	"encoding/json"
	"encoding/base64"
	"fmt"
	"github.com/nu7hatch/gouuid"
	"io"
	"crypto/sha256"
	"crypto/hmac"
)

const HMACMESSAGE string = "S3cr3tp@ss"

var cookieHis = make(map[string]string)

type Person struct {
	Name string
	Age  int
}

func encodeJson(user Person) string{
	userJson, _ := json.Marshal(user)
	fmt.Println(string(userJson))
	return base64.StdEncoding.EncodeToString(userJson)
}

func createHMAC(message string) string{
	h := hmac.New(sha256.New, []byte(HMACMESSAGE))
	io.WriteString(h, message)
	return fmt.Sprintf("%x", h.Sum(nil))
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
	for key, value := range cookieHis {
		fmt.Println("Key:", key, "Value:", value)
		cookie, err := req.Cookie(key)
		if err == nil {
			fmt.Println("Cookie:" + cookie.String())
			// Logging the possible errors
			if value != createHMAC(cookie.Value) {
				fmt.Println("Cookie has been changed: [" + key + "]")
				return
			}
		} else {
			fmt.Println(err)
		}
	}
	// Generating a new ID
	id, err := uuid.NewV4()
	if(err == nil) {
		cookie := &http.Cookie{
			Name: "sessionID",
			Value: id.String(),
			HttpOnly:true,
			//Secure: true,
		}
		cookieHis["sessionID"] = createHMAC(id.String())
		http.SetCookie(res, cookie)
	}
	//User info
	name := req.FormValue("name")
	age, _ := strconv.Atoi(req.FormValue("age"))
	p := Person{Name: name, Age: age}
	if req.Method == "POST" {
		cookie, err := req.Cookie("userinfo")
		if err == http.ErrNoCookie{

			cookie = &http.Cookie{
				Name: "userinfo",
				Value: encodeJson(p),
				HttpOnly:true,
				//Secure: true,
			}
		}
		// Setting the cookie on the response back to the client
		cookieHis["userinfo"] = encodeJson(p)
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

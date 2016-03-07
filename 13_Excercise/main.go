package main

import (
	"net/http"
	"strconv"
	"fmt"
)


func cookieHandler(res http.ResponseWriter, req *http.Request){
	cookie, err := req.Cookie("count-cookie")
	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  "count-cookie",
			Value: "0",
		}
	}
	count,_ := strconv.Atoi(cookie.Value)
	count++;
	http.SetCookie(res, &http.Cookie{
		Name:  "count-cookie",
		Value: strconv.Itoa(count),
	})
	fmt.Fprintf(res, "<h1>Check counter-cookie by developer tool to assert its value is:%s</h1>",
		strconv.Itoa(count))
}

func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", cookieHandler)
	http.ListenAndServe(":8080", nil)
}

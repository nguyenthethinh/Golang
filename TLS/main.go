package main

import (
"net/http"
"log"
)

func main() {
	http.HandleFunc("/", index)
	log.Println("Redirect to port 10443")
	go http.ListenAndServe(":8080", http.RedirectHandler("https://localhost:10443", http.StatusMovedPermanently))
	err := http.ListenAndServeTLS(":10443", "cert.pem", "key.pem", nil)
	if (err != nil) {
		log.Fatal(err)
	}
}

func index(res http.ResponseWriter, req *http.Request){
	res.Header().Set("Content-Type", "text/plain")
	res.Write([]byte("This is testing TLS server"))
}

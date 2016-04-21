package main
import (
	"net/http"
	"html/template"
	"log"
	"github.com/nu7hatch/gouuid"
	"encoding/json"
	"encoding/base64"
	"crypto/hmac"
	"crypto/sha256"
	"io"
	"fmt"
)

var template1 *template.Template
type Person struct {
	Uuid, Name, Hmac string
	Loggedin         bool
}
func init() {
	var err error
	template1, err = template.ParseFiles("index.html")
	if(err != nil){
		log.Println("Error: ", err)
	}
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/logout", logout)
	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	cookie, err1 := req.Cookie("user-info")
	p := Person{};
	if(err1 == http.ErrNoCookie){
		uuid, _ := uuid.NewV4()
		p = Person{
			Uuid: uuid.String(),
			Hmac: getCode(uuid.String()),
		}
		encodedUser := encodeJson(p)
		cookie = &http.Cookie{
			Name: "user-info",
			Value: encodedUser,
			HttpOnly: true,
			//		Secure: true,
		}
		http.SetCookie(res, cookie)
		p.Loggedin = true
	}
	if req.Method == "POST" {
		log.Println("Posting request....")
		name := req.FormValue("name")
		p = Person{
			Name: name,
		}
		var e error
		cookie, e = req.Cookie("user-info")
		if (e != nil) {
			log.Println("Can't get cookie")
			return;
		}
		cookie.Value = updateCookie(p, req, cookie)
		p = decodeJson(cookie)
	}
	template1.Execute(res, p)
}

func logout(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		log.Println("Posting logout request....")
		p := Person{
			Loggedin: false,
		}
		template1.Execute(res, p)
	}
}

func updateCookie(person Person, req *http.Request, cookie *http.Cookie) string {
	decodedPerson := decodeJson(cookie)
	if decodedPerson.Loggedin == false {
		log.Println("Authentication failed!")
		return encodeJson(person)
	}
	person.Uuid = decodedPerson.Uuid
	person.Hmac = getCode(person.Uuid + person.Name)
	return encodeJson(person)
}

func getCode(data string) string {
	h := hmac.New(sha256.New, []byte("key"))
	io.WriteString(h, data)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func encodeJson(user Person) string {
	jsonUser, errJsonMarshalError := json.Marshal(user)
	if errJsonMarshalError != nil {
		log.Println("Error: ", errJsonMarshalError)
	}
	return base64.StdEncoding.EncodeToString(jsonUser)
}

func decodeJson(cookie *http.Cookie) Person {
	log.Println("Cookie", cookie.Value)
	decode, err := base64.StdEncoding.DecodeString(cookie.Value)
	if err != nil {
		log.Println("Error: ", err)
		var person Person
		person.Loggedin = false
		return person
	}
	var person Person
	json.Unmarshal(decode,&person)
	if person.Hmac == getCode(person.Uuid + person.Name) {
		log.Println("Cookie have not been changed!")
		person.Loggedin = true
		return person
	}
	log.Println("Cookie changed!")
	person.Loggedin = false
	return person
}
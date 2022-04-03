package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"text/template"
)

var TheNumber int
var CurrentUser UserDetalies = UserDetalies{"aaa", "", false}

type UserDetalies struct {
	Login    string
	Password string
	Success  bool
}

func main() {
	initTheNumber()
	StartWebApp()
}

func initTheNumber() {
	TheNumber = rand.Intn(100)
}

func StartWebApp() {
	port := os.Getenv("PORT")
	http.HandleFunc("/", index)
	http.ListenAndServe(port, nil)
}

func index(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		CurrentUser.Login = request.PostFormValue("login")
		CurrentUser.Password = request.PostFormValue("password")
		Login(&CurrentUser)
		fmt.Println(CurrentUser.Success)

		t, _ := template.ParseFiles("templates/index.html")
		t.ExecuteTemplate(writer, "index", CurrentUser)
	} else {
		t, _ := template.ParseFiles("templates/index.html")
		t.ExecuteTemplate(writer, "index", CurrentUser)
	}
}

func Login(CtUser *UserDetalies) {
	if (*CtUser).Login == "user" && (*CtUser).Password == "pass" {
		(*CtUser).Success = true
	} else {
		(*CtUser).Success = false
	}

}

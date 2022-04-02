package main

import (
	"math/rand"
	"net/http"
)

var TheNumber int

func main() {
	initTheNumber()
	StartWebApp()
}

func initTheNumber() {
	TheNumber = rand.Intn(100)
}

func StartWebApp() {
	http.HandleFunc("/", index)
}

func index(writer http.ResponseWriter, request *http.Request) {

}

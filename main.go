package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("views/html/index.html")
		t.Execute(w, nil)
	}
}

func main() {
	http.HandleFunc("/", index)              // setting router rule
	err := http.ListenAndServe(":9090", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

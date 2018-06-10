package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	"www/src"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("views/html/about.html")
		t.Execute(w, nil)
	}
}

func projects(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		user, err := github.GetUser("DJBrunelle")
		if err != nil {
			println("Unable to get user")
		}
		t, _ := template.ParseFiles("views/html/projects.html")
		t.Execute(w, user)
	}
}

func hireMe(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("views/html/hire_me.html")
		t.Execute(w, nil)
	}
}

func main() {
	http.HandleFunc("/", index)              // setting router rule
	http.HandleFunc("/about", index)         // setting router rule
	http.HandleFunc("/projects", projects)   // setting router rule
	http.HandleFunc("/hire_me", hireMe)      // setting router rule
	err := http.ListenAndServe(":5000", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

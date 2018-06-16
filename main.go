package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
	"www/src"
)

//Index route handler
func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("views/html/about.html")
		t.Execute(w, nil)
	}
}

//Projects route handler
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

//HireMe route handler
func hireMe(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("views/html/hire_me.html")
		t.Execute(w, nil)
	}
}

//Set up routes and open port to listen for requests
func main() {
	http.HandleFunc("/", index)                            // setting router rule
	http.HandleFunc("/about", index)                       // setting router rule
	http.HandleFunc("/projects", projects)                 // setting router rule
	http.HandleFunc("/hire_me", hireMe)                    // setting router rule
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

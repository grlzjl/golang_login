package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("views/login.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
		if r.FormValue("username") == "zhangsan" && r.FormValue("password") == "123" {
			http.HandleFunc("/success", success)
			http.Redirect(w, r, "/success", http.StatusFound)

		} else {
			http.HandleFunc("/fail", fail)
			http.Redirect(w, r, "/fail", http.StatusFound)
		}
	}
}

func success(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("views/success.html")
		t.Execute(w, nil)
	}
	// fmt.Fprintf(w, "Hello world, this is my first page!")
	//w.Header().Set("Content-Type", "text/html")
	//html := `<doctype html> <html> <head> <title>Page Message</title> </head> <body> <p> <a href="/Index">Index</a> | <a href="/Welcome">Welcome</a> </p> </body> </html>`
	//fmt.Fprintln(w, html)
}
func fail(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("views/fail.html")
		t.Execute(w, nil)
	}
	// fmt.Fprintf(w, "Hello world, this is my first page!")
	//w.Header().Set("Content-Type", "text/html")
	//html := `<doctype html> <html> <head> <title>Page Welcome</title> </head> <body> <p> <a href="/Index">Index</a> | <a href="/Message">Message</a> </p> </body> </html>`
	//fmt.Fprintln(w, html)
}


func main() {
	http.Handle("/static", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/login", login)

	err := http.ListenAndServe(":8083", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

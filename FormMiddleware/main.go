package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

type ContactDetails struct {
	Email   string
	Subject string
	Message string
}
type TemplateData struct {
	Success bool
	Details ContactDetails
}

func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		f(w, r)
	}
}

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "foo")
}

func bar(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "bar")
}

func middlewareOne(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Print("Executing middlewareOne")
		next.ServeHTTP(w, r)
		log.Print("Executing middlewareOne again")
	}
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func main() {
	tmpl := template.Must(template.ParseFiles("assets/templates/layout.html"))
	formtmpl := template.Must(template.ParseFiles("assets/templates/form.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := TodoPageData{
			PageTitle: "My Todo List",
			Todos: []Todo{
				{Title: "Go For A Walk", Done: true},
				{Title: "Buy Milk", Done: false},
				{Title: "Learn New Course", Done: true},
			},
		}
		tmpl.Execute(w, data)
	})

	http.HandleFunc("/form", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			formtmpl.Execute(w, nil)
			return
		}
		details := ContactDetails{
			Email:   r.FormValue("email"),
			Subject: r.FormValue("subject"),
			Message: r.FormValue("message"),
		}
		deta := details
		data := TemplateData{
			Success: true,
			Details: deta,
		}
		fmt.Println(deta)

		formtmpl.Execute(w, data)
	})
	http.HandleFunc("/foo", logging(foo))
	http.HandleFunc("/bar", logging(bar))
	http.HandleFunc("/boo", middlewareOne(handleHome))
	http.ListenAndServe(":8080", nil)
}

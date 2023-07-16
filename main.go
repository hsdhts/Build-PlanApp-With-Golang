package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// Todo struct
type Todo struct {
	ID    int
	Title string
	Done  bool
}

// In-memory todos
var todos []Todo

func main() {
	http.HandleFunc("/", todoHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":8080", nil)
}

func todoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		title := r.FormValue("title")
		addTodo(title)
	}

	renderTemplate(w, "index.html", todos)
}

func addTodo(title string) {
	todo := Todo{
		ID:    len(todos) + 1,
		Title: title,
		Done:  false,
	}
	todos = append(todos, todo)
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	tmplFile := fmt.Sprintf("templates/%s", tmpl)
	t, err := template.ParseFiles(tmplFile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

package main

import (
	"log"
	"net/http"
	"text/template"
)

var tmp *template.Template

type Todo struct {
	Item string
	Done bool
}
type PageData struct {
	Title string
	Todos []Todo
}

func ToDo(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title: "To Do List",
		Todos: []Todo{
			{Item: "Install Go", Done: true},
			{Item: "Learn Go", Done: false},
			{Item: "like this", Done: false},
		},
	}
	tmp.Execute(w, data)
}
func main() {
	mux := http.NewServeMux()
	tmp = template.Must(template.ParseFiles("template/index.gohtml"))
	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	mux.HandleFunc("/todo", ToDo)
	log.Fatal(http.ListenAndServe(":5000", mux))
}

package main

import (
	"net/http"
	"golang-web-dev/045-code-organization/03_multiple-packages/books"
)

func main() {
	http.HandleFunc("/", books.Index)
	http.HandleFunc("/books", books.Index)
	http.HandleFunc("/books/show", books.Show)
	http.HandleFunc("/books/create", books.Create)
	http.HandleFunc("/books/create/process", books.CreateProcess)
	http.HandleFunc("/books/update", books.Update)
	http.HandleFunc("/books/update/process", books.UpdateProcess)
	http.HandleFunc("/books/delete/process", books.DeleteProcess)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request)  {
	http.Redirect(w, r, "/books", http.StatusSeeOther)
}
package main

import (
	r "./report"
	"net/http"
	"log"
)

func main() {
    handleRequests()
}

func handleRequests() {
    log.Println("Server started on: http://localhost:8080")
	http.HandleFunc("/", r.Index)
	http.HandleFunc("/show", r.Show)
    http.HandleFunc("/new", r.New)
   	http.HandleFunc("/edit", r.Edit)
    http.HandleFunc("/insert", r.Insert)
    http.HandleFunc("/update", r.Update)
    http.HandleFunc("/delete", r.Delete)
    http.HandleFunc("/search", r.Search)
    
    http.ListenAndServe(":8080", nil)
}

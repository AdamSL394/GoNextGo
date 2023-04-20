package main

import (
	"fmt"
	"encoding/json"
	"net/http"
)

type Article struct {
	Id     string
	Title  string
	Desc    string
	Content string
}

var Articles []Article

func main() {

	Articles = []Article{
		{Id: "1", Title: "First article", Desc: "Title of this fine article", Content: "Content for this fine article"},
		{Id: "2", Title: "Second article", Desc: "Title of this majestic article", Content: "Content for this majestic article"},
	}

	http.HandleFunc("/articles", foo)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		fmt.Fprintf(w, "Hello Adam, you've requested: %s\n", r.URL.Path);
	})

	http.ListenAndServe(":8002", nil)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	}

func foo(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		js, err := json.Marshal(Articles)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
}
package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Whoa, Go is a bad boy!")
}

func newsAggHandler(w http.ResponseWriter, r *http.Request) {
	p := NewsAggPage{Title: "Amazing News aggregator", News: "Some news"}
	t, _ := template.ParseFiles("basictemplating.html")
	t.Execute(w, p)
}

type NewsAggPage struct {
	Title string
	News  string
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/agg/", newsAggHandler)
	http.ListenAndServe(":8000", nil)
}

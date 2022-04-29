package main

import (
	"fmt"
	"net/http"
	"log"
	"encoding/json"
)

type Article struct {
	Title string `json:"title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}
type Articles []Article

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello world!")
}

func allArticles(w http.ResponseWriter, r *http.Request){
	article := Articles{
		Article{ Title: "Hello wolrd", Desc: "This article is about hello world", Content: "World is proximately 4 milion year olds i guess :)" },
	}

	fmt.Println("Hit Articles endpoint!")
	json.NewEncoder(w).Encode(article)
}

func handleRequests(){
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequests()
}
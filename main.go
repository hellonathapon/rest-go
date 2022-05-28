package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Test struct {
	Title string `json:"title"`
	Sub string `json:"sub"`
}

type Article struct {
	Title string `json:"title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}
type Articles []Article // array of struct [{}]

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello world!")
}

func testPostReq(w http.ResponseWriter, r *http.Request){
	// expect json format to be passed to this route handler
	// curl -X POST -d "{\"Title\":\"Works\"}" http://localhost:8081/test
	decoder := json.NewDecoder(r.Body)
	var t Test
	err := decoder.Decode(&t) // pass address of t
	if err != nil {
		panic(err)
	}
	fmt.Println(t)
	// fmt.Fprintf(w, t.Title) // response string
	json.NewEncoder(w).Encode(t) // response json
}


func allArticles(w http.ResponseWriter, r *http.Request){
	article := Articles{
		Article{ Title: "Hello wolrd", Desc: "This article is about hello world", Content: "World is proximately 4 milion year olds i guess :)" },
	}

	fmt.Fprintf(w, "Hit Articles endpoint!") // Fprintf is actually a reponse as well
	json.NewEncoder(w).Encode(article)
}

func handleRequests(){
	myRouter := mux.NewRouter().StrictSlash(true)
	/**
	 * * Map all routes
	 */
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", allArticles).Methods("GET")
	myRouter.HandleFunc("/test", testPostReq).Methods("POST")

	/**
	 * * log.Fatal is essentially logging function 
	 * * works when fails to starting the server. 
	 */
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	handleRequests()
}
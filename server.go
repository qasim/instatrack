package main

import (
	"log"
	"net/http"
	"os"
)

type item struct {
	AuthorName  string `json:"authorName"`
	AuthorImage string `json:"authorImage"`
	Image       string `json:"image"`
	Text        string `json:"text"`
	Date        string `json:"date"`
}

func handleItems(w http.ResponseWriter, r *http.Request) {
	return
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	http.HandleFunc("/api/items", handleItems)
	http.Handle("/", http.FileServer(http.Dir("./www")))
	log.Println("Server running at http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

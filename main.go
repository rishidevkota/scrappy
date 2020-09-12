package main

import (
	"log"
	"net/http"

	"insta_graph/handler"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := r.Cookie("authtoken")
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		http.Redirect(w, r, "/home", http.StatusFound)
	})
	http.HandleFunc("/login", handler.Login)
	http.HandleFunc("/callback", handler.Callback)
	http.HandleFunc("/home", handler.Home)
	http.ListenAndServe(":8080", nil)
}

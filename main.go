package main

import (
	"blog/handlers"
	"blog/storage"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	//Starting server
	connStr := fmt.Sprintf(
		"host=localhost port=5432 user=postgres password=%s dbname=blog sslmode=disable",
		os.Getenv("DB_PASSWORD"),
	)
	db, err := storage.NewDB(connStr)
	h := handlers.NewHandler(db)
	if err != nil {
		log.Fatal(err)
	}
	storage.CreateTable(db)

	fmt.Println("Сервер запущен на порту 8080")
	http.HandleFunc("/posts", h.HandlePosts)
	http.HandleFunc("/post", h.HandlePost)
	http.HandleFunc("/comments", h.CommentHandler)
	http.ListenAndServe(":8080", nil)
}

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
	connStr := fmt.Sprintf(
		"host=localhost port=5432 user=postgres password=%s dbname=blog sslmode=disable",
		os.Getenv("DB_PASSWORD"),
	)
	db, err := storage.NewDB(connStr)
	if err != nil {
		log.Fatal(err)
	}
	if err = storage.CreateTable(db); err != nil {
		log.Fatal(err)
	}
	if err = storage.CreateCommentsTable(db); err != nil {
		log.Fatal(err)
	}

	h := handlers.NewHandler(db)
	http.HandleFunc("/posts", h.HandlePosts)
	http.HandleFunc("/post", h.HandlePost)
	http.HandleFunc("/comments", h.CommentHandler)

	fmt.Println("Сервер запущен на порту 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

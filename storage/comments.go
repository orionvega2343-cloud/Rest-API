package storage

import (
	"blog/models"
	"database/sql"
)

func CreateCommentsTable(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS comments (
     id SERIAL PRIMARY KEY,
     text TEXT NOT NULL,
     author TEXT NOT NULL,
     post_id INTEGER NOT NULL)`)
	if err != nil {
		return err
	}
	return nil
}
func GetComments(db *sql.DB, postID int) ([]models.Comment, error) {
	rows, err := db.Query("SELECT id, text, author, post_id FROM comments WHERE post_id = $1", postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var comments []models.Comment
	for rows.Next() {
		var p models.Comment
		rows.Scan(&p.ID, &p.Text, &p.Author, &p.PostID)
		comments = append(comments, p)
	}
	return comments, nil
}

func CreateComments(db *sql.DB, p models.Comment) (models.Comment, error) {
	rows := db.QueryRow("INSERT INTO comments (text,author,post_id) VALUES ($1, $2, $3) RETURNING id", p.Text, p.Author, p.PostID)
	err := rows.Scan(&p.ID)
	if err != nil {
		return models.Comment{}, err
	}
	return p, nil
}

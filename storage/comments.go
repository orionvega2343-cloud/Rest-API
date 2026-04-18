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
	return err
}

func GetComments(db *sql.DB, postID int) ([]models.Comment, error) {
	rows, err := db.Query("SELECT id, text, author, post_id FROM comments WHERE post_id = $1", postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var comments []models.Comment
	for rows.Next() {
		var c models.Comment
		if err = rows.Scan(&c.ID, &c.Text, &c.Author, &c.PostID); err != nil {
			return nil, err
		}
		comments = append(comments, c)
	}
	return comments, nil
}

func CreateComments(db *sql.DB, c models.Comment) (models.Comment, error) {
	err := db.QueryRow(
		"INSERT INTO comments (text, author, post_id) VALUES ($1, $2, $3) RETURNING id",
		c.Text, c.Author, c.PostID,
	).Scan(&c.ID)
	if err != nil {
		return models.Comment{}, err
	}
	return c, nil
}

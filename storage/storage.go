package storage

import (
	"blog/models"
	"database/sql"
	"errors"

	_ "github.com/lib/pq"
)

var ErrNotFound = errors.New("not found")

func NewDB(connStr string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connStr) //Создаем бд

	if err != nil {
		return nil, err
	}

	err = db.Ping() //Запускаем бд
	if err != nil {
		return nil, err
	}
	return db, nil
}

func CreateTable(db *sql.DB) error { //Создаем таблицу
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS posts (
    					id SERIAL PRIMARY KEY, 
						title TEXT NOT NULL,
						text TEXT NOT NULL,
						author TEXT NOT NULL,
						date TEXT NOT NULL)`)
	if err != nil {
		return err
	}
	return nil
}

func GetAll(db *sql.DB) ([]models.Post, error) {
	rows, err := db.Query("SELECT id, title, text, author, date FROM posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var posts []models.Post
	for rows.Next() {
		var p models.Post
		rows.Scan(&p.ID, &p.Title, &p.Text, &p.Author, &p.Date)
		posts = append(posts, p)

	}

	return posts, nil
}

func Create(db *sql.DB, p models.Post) (models.Post, error) {
	row := db.QueryRow("INSERT INTO posts (title, text, author, date) VALUES ($1, $2, $3, $4) RETURNING id", p.Title, p.Text, p.Author, p.Date)
	err := row.Scan(&p.ID)
	if err != nil {
		return models.Post{}, err
	}
	return p, nil

}

func Update(db *sql.DB, id int, p models.Post) (models.Post, error) {
	result, err := db.Exec(
		"UPDATE posts SET title=$1, text=$2, author=$3, date=$4 WHERE id=$5",
		p.Title, p.Text, p.Author, p.Date, id,
	)
	if err != nil {
		return models.Post{}, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return models.Post{}, err
	}
	if rowsAffected == 0 {
		return models.Post{}, ErrNotFound
	}

	p.ID = id
	return p, nil
}

func Delete(db *sql.DB, id int) error {
	result, err := db.Exec("DELETE FROM posts WHERE id = $1", id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrNotFound
	}
	return nil
}

func GetByID(db *sql.DB, id int) (models.Post, error) {
	var p models.Post
	err := db.QueryRow(
		"SELECT id, title, text, author, date FROM posts WHERE id = $1", id,
	).Scan(&p.ID, &p.Title, &p.Text, &p.Author, &p.Date)

	if err == sql.ErrNoRows {
		return models.Post{}, ErrNotFound
	}
	if err != nil {
		return models.Post{}, err
	}
	return p, nil
}

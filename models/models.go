package models

type Post struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Text   string `json:"text"`
	Author string `json:"author"`
	Date   string `json:"date"`
}

type Comment struct {
	ID     int    `json:"id"`
	Text   string `json:"text"`
	Author string `json:"author"`
	PostID int    `json:"post_id"`
}

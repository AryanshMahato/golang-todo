package model

import "database/sql"

var DB *sql.DB

type Todo struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

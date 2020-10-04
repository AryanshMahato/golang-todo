package model

import "database/sql"

var DB *sql.DB

type Todo struct {
	ID    string `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
}

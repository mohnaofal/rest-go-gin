package models

import "time"

type Article struct {
	ID      uint      `json:"id" db:"id"`
	Author  string    `json:"author" db:"author"`
	Title   string    `json:"title" db:"title"`
	Body    string    `json:"body" db:"body"`
	Created time.Time `json:"created" db:"created"`
}

type ArticleParams struct {
	Query  string `json:"query"`
	Author string `json:"author"`
}

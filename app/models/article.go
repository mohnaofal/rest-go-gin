package models

import "time"

type Article struct {
	ID      uint      `json:"id" db:"id"`
	Author  string    `json:"author" db:"author" form:"author" binding:"required"`
	Title   string    `json:"title" db:"title" form:"title" binding:"required"`
	Body    string    `json:"body" db:"body" form:"body" binding:"required"`
	Created time.Time `json:"created" db:"created"`
}

type ArticleParams struct {
	Query  string `json:"query"`
	Author string `json:"author"`
}

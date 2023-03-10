package models

import (
	"gorm.io/gorm"
)

/*
gorm.model will add properties such as ID, CreatedAt, UpdatedAt and DeletedAt
*/
type Entry struct {
	gorm.Model
	Rank  int    `json:"rank"`
	Title string `json:"title"`
	URL   string `json:"url"`
}

type Post struct {
	Entry
	Content    string `json:"content"`
	Score      string `json:"score"`
	Comments   int    `json:"comments"`
	PostTime   string `json:"postTime,omitempty"`
	AuthorUrl  string `json:"authorUrl,omitempty"`
	AuthorName string `json:"authorName,omitempty"`
}

package models

import (
	"gorm.io/gorm"
)

type Entry struct {
	gorm.Model        // gorm.model will add properties such as ID, CreatedAt, UpdatedAt and DeletedAt
	Rank       int    `json:"rank"`  // The rank of the entry
	Title      string `json:"title"` // The title of the entry
	URL        string `json:"url"`   // The URL of the entry
}

type Post struct {
	Entry             // A Post is an Entry, so it includes all the properties of an Entry
	Content    string `json:"content"`              // The content of the post
	Score      string `json:"score"`                // The score of the post
	Comments   int    `json:"comments"`             // The number of comments on the post
	PostTime   string `json:"postTime,omitempty"`   // The time that the post was made (optional)
	AuthorUrl  string `json:"authorUrl,omitempty"`  // The URL of the author of the post (optional)
	AuthorName string `json:"authorName,omitempty"` // The name of the author of the post (optional)
}

package models

// Tag ...
type Tag struct {
	TagID int    `json:"tagID"`
	Title string `json:"title"`
}

// Tags ...
type Tags []Tag

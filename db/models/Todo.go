package models

import (
	// "time"
)

type Todo struct {
	Id int64 `json:"id"`
	Title string `json:"title"`
	Body string `json:"body"`
	CreatedAt string `json:"createdAt"`
}
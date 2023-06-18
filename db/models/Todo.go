package models

import (
	// "time"
)

type Todo struct {
	Id int64 `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	CreatedAt string `json:"createdAt"`
}
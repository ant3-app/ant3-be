package models

import (
	"time"
)

type TableQueue struct {
	Id string `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Name string `json:"name"`
	Number int `json:"number"`
}
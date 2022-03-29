package model

import (
	"fmt"

	"gorm.io/gorm"
)

type Authors struct {
	Authors []Author `json:"Authors"`
}

// Create Author type
type Author struct {
	gorm.Model
	ID   int    `json:"AuthorID"`
	Name string `json:"Name"`
}

func (a *Author) ToString() string {

	return fmt.Sprintf("ID: %v Name: %s", a.ID, a.Name)
}

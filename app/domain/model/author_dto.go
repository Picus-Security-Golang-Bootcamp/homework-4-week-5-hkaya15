package model

import (
	"fmt"

	"gorm.io/gorm"
)

type AuthorsWithBook struct {
	gorm.Model
	ID       int    `gorm:"foreignKey:ID"`
	Name     string `gorm:"foreignKey:name"`
	Page     string `gorm:"foreignKey:page"`
	BookName string `gorm:"foreignKey:bookname"`
}

func (a *AuthorsWithBook) ToString() string {

	return fmt.Sprintf("ID: %v Name: %s Page: %v BookName: %s", a.ID, a.Name,a.Page,a.BookName)
}

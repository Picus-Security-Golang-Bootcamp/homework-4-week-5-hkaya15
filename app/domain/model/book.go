package model

import (
	"fmt"

	"gorm.io/gorm"
)

type Books struct {
	Books []Book `json:"Books"`
}

// Create Book type
type Book struct {
	gorm.Model
	BookID     string  `json:"BookID" gorm:"not_null"`
	BookName   string  `json:"Name"`
	Page       uint    `json:"Page"`
	StockCount uint    `json:"StockCount"`
	Price      float64 `json:"Price" gorm:"type:numeric(10,2)"`
	StockID    int     `json:"StockID"`
	ISBN       int     `json:"ISBN"`
	AuthorID   uint    `json:"AuthorID"`
	Authors    Author  `gorm:"foreignKey:ID; references:AuthorID"`
}

func (b *Book) ToString() string {
	return fmt.Sprintf("ID : %v, Book Name : %s, Page : %v, StockCount : %v,Price : %v, StockID : %v, ISBN : %v, AuthorID: %v, Author Name: %v,CreatedAt : %s", b.BookID, b.BookName, b.Page, b.StockCount, b.Price, b.StockID, b.ISBN, b.AuthorID, b.Authors.Name, b.CreatedAt.Format("2006-01-02 15:04:05"))
}

func (b *Book) BeforeDelete(tx *gorm.DB) (err error) {
	fmt.Printf("Book (%s) deleting...\n", b.BookName)
	return nil
}

func (b *Book) AfterDelete(tx *gorm.DB) (err error) {
	fmt.Printf("Book (%s) deleted!\n", b.BookName)
	return nil
}

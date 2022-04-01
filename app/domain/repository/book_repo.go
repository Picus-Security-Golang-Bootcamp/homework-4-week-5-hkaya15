package repository

import (
	"encoding/json"
	"io/ioutil"
	"os"

	. "github.com/hkaya15/PicusSecurity/Week_5_Homework/app/domain/model"
	. "github.com/hkaya15/PicusSecurity/Week_5_Homework/base/logs"
	"gorm.io/gorm"
)

var log = NewBuiltinLogger()

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{db: db}
}

// Migrations helps the auto-migrate
func (b *BookRepository) Migrations() {
	b.db.AutoMigrate(&Book{})
}

// InsertData helps the insert data
func (b *BookRepository) InsertData() error {
	bookList, err := getAllBooksFromJSON()
	if err != nil {
		log.Logger.Println(err.Error())
		return err
	}

	for _, v := range bookList.Books {
		b.db.Where(Book{BookID: v.BookID}).Attrs(Book{BookID: v.BookID, BookName: v.BookName, Page: v.Page, StockCount: v.StockCount, Price: v.Price, StockID: v.StockID, ISBN: v.ISBN, AuthorID: v.AuthorID}).FirstOrCreate(&v)
	}
	return nil
}

// CreateBook returns the book which created
func (b *BookRepository) CreateBook(book *Book) (Book, error) {
	result := b.db.Create(&book)
	if result.Error != nil {
		log.Logger.Println(result.Error.Error())
		return *book, result.Error
	}
	return *book, nil
}

// UpdateBook returns the updated book
func (b *BookRepository) UpdateBook(book *Book) (Book, error) {
	result := b.db.Model(&book).Where("book_id = ?", book.BookID)
	if result.Error != nil {
		log.Logger.Println(result.Error.Error())
		return *book, result.Error
	}

	result.Save(&book)
	return *book, nil
}

// GetByID returns an book by id
func (b *BookRepository) GetByID(id string) (*Book, error) {
	var book Book

	result := b.db.Where("book_id = ?", id).First(&book)
	if result.Error != nil {
		log.Logger.Println(result.Error.Error())
		return &book, result.Error
	}
	return &book, nil
}

// FindAll returns the book list
func (b *BookRepository) FindAll() ([]Book, error) {
	var books []Book
	result := b.db.Unscoped().Where("deleted_at IS NULL").Find(&books)
	if result.Error != nil {
		log.Logger.Println(result.Error.Error())
		return nil, result.Error
	}
	return books, nil
}

// GetBooksWithAuthor returns the booklist with authors
func (b *BookRepository) GetBooksWithAuthor() ([]Book, error) {
	var books []Book

	result := b.db.Preload("Authors").Find(&books)
	if result.Error != nil {
		log.Logger.Println(result.Error.Error())
		return nil, result.Error
	}

	return books, nil
}

// SearchName returns the book list by book name regarding of the contains & non-case sensitive
func (b *BookRepository) SearchByName(name string) ([]Book, error) {
	var books []Book
	result := b.db.Where("book_name ILIKE ? ", "%"+name+"%").Find(&books)
	if result.Error != nil {
		log.Logger.Println(result.Error.Error())
		return nil, result.Error
	}
	return books, nil
}

// BuyByID returns the book that buy by book ıd & book count if it causes negative value (uint)
func (b *BookRepository) BuyByID(id string, count uint) (Book, error) {
	var book Book
	b.db.Model(&book).Where(Book{BookID: id}).Find(&book)
	book.StockCount = book.StockCount - count
	result := b.db.Save(&book)
	if result.Error != nil {
		log.Logger.Println(result.Error.Error())
		return book, result.Error
	}
	return book, nil
}

// DeleteByID returns the book that deleted by book id. It just update Deleted_At on DB (soft delete)
func (b *BookRepository) DeleteByID(id string) error {
	result := b.db.Unscoped().Where("book_id = ?", id).Delete(&Book{})
	if result.Error != nil {
		log.Logger.Println(result.Error.Error())
		return result.Error
	}
	return nil
}

// Pagination returns the books regarding of offset & limit value
func (b *BookRepository) Pagination(offSet int, pageSize int) ([]Book, error) {
	var book []Book
	result := b.db.Offset(offSet).Limit(pageSize).Find(&book)
	if result.Error != nil {
		log.Logger.Println(result.Error.Error())
		return book, result.Error
	}
	return book, nil
}

// getAllBooksFromJSON returns the book list that readed by json
func getAllBooksFromJSON() (*Books, error) {
	var books Books
	jsonFile, err := os.Open(os.Getenv("BOOK_JSON"))
	if err != nil {
		log.Logger.Println(err.Error())
		return nil, err
	}

	byteVal, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteVal, &books)

	defer jsonFile.Close()

	return &books, nil
}

package repository

import (
	"encoding/json"
	"io/ioutil"
	"os"

	. "github.com/hkaya15/PicusSecurity/Week_5_Homework/app/domain/model"
	"gorm.io/gorm"
)

type AuthorRepository struct {
	db *gorm.DB
}



func NewAuthorRepository(db *gorm.DB) *AuthorRepository {
	return &AuthorRepository{db: db}
}

// Migrations helps the auto-migrate
func (a *AuthorRepository) Migrations() {
	a.db.AutoMigrate(&Author{})
}

// InsertData helps the insert data
func (c *AuthorRepository) InsertData() error {
	authorsList, err := getAllAuthorsFromJSON()
	if err != nil {
		log.Logger.Fatalln(err.Error())
		return err
	}

	for _, v := range authorsList.Authors {
		c.db.Where(Author{ID: v.ID}).Attrs(Author{ID: v.ID, Name: v.Name}).FirstOrCreate(&v)
	}
	return nil
}

// GetByID returns an author by id
func (a *AuthorRepository) GetByID(id int) (Author, error) {
	var author Author

	result := a.db.Where("id = ?", id).First(&author)
	if result.Error != nil {
		log.Logger.Println(result.Error.Error())
		return Author{}, result.Error
	}
	return author, nil
}

// SearchByName returns Authors list by name
func (a *AuthorRepository) SearchByName(name string) ([]Author, error) {
	var authors []Author
	result := a.db.Where("Name ILIKE ? ", "%"+name+"%").Find(&authors)
	if result.Error != nil {
		log.Logger.Println(result.Error.Error())
		return nil, result.Error
	}
	return authors, nil
}

// GetAuthorsWithBooks return DTO combined with book names
func (a *AuthorRepository) GetAuthorsWithBooks() ([]AuthorsWithBook, error) {
	var authors []AuthorsWithBook
	result := a.db.Table("authors").Select("authors.id,authors.name,books.page,books.book_name").Joins("Inner join books on books.author_id = authors.id").Find(&authors)
	if result.Error != nil {
		log.Logger.Println(result.Error.Error())
		return []AuthorsWithBook{}, result.Error
	}
	return authors, nil
}

// UpdateAuthorName updates author name
func (a *AuthorRepository) UpdateAuthorName(author *Author) (Author, error) {
	result := a.db.Model(&author).Where("id = ?", author.ID).Update("name", author.Name)
	if result.Error != nil {
		log.Logger.Println(result.Error.Error())
		return *author, result.Error
	}
	return *author, nil
}

// FindAll returns all author list
func (a *AuthorRepository) FindAll() ([]Author, error) {
	var authors []Author
	result := a.db.Find(&authors)
	if result.Error != nil {
		log.Logger.Println(result.Error.Error())
		return authors, result.Error
	}
	return authors, nil
}

// getAllAuthorsFromJSON convert JSON to Author slice
func getAllAuthorsFromJSON() (*Authors, error) {

	var authors Authors
	jsonFile, err := os.Open(os.Getenv("AUTHOR_JSON"))
	if err != nil {
		log.Logger.Fatalln(err.Error())
		return nil, err
	}

	byteVal, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteVal, &authors)

	defer jsonFile.Close()

	return &authors, nil
}

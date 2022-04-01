package server

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	//. "github.com/hkaya15/PicusSecurity/Week_5_Homework/api/models/api"

	// . "github.com/hkaya15/PicusSecurity/Week_5_Homework/api/server/utils/helpers"
	model "github.com/hkaya15/PicusSecurity/Week_5_Homework/app/domain/model"
	. "github.com/hkaya15/PicusSecurity/Week_5_Homework/app/domain/repository"
)

type BookHelper struct {
	Repo *BookRepository
}

func NewBookHandler(b *BookRepository) BookHelper {
	return BookHelper{Repo: b}
}

// Create Book
func (bk BookHelper) CreateBook(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var book model.Book
	json.Unmarshal(body, &book)

	result, error := bk.Repo.CreateBook(&book)
	if error != nil {

		http.Error(w, error.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&result)

}

// UpdateBook
func (bk BookHelper) UpdateBook(w http.ResponseWriter, r *http.Request) {
	var book model.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	bk.Repo.UpdateBook(&book)

	if err := json.NewEncoder(w).Encode(&book); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

// FindBookById
func (bk BookHelper) FindBookById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	result, error := bk.Repo.GetByID(id)

	if error != nil {
		http.Error(w, error.Error(), http.StatusInternalServerError)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&result)

}

//DeleteBookById

func (bk BookHelper) DeleteBookById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	error := bk.Repo.DeleteByID(id)

	if error != nil {
		http.Error(w, error.Error(), http.StatusInternalServerError)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Success"))

}


// FindBookByName
func (bk BookHelper) FindBookByName(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	name := vars["name"]
	result, error := bk.Repo.SearchByName(name)
	if error != nil {
		http.Error(w, error.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&result)

}

/////

func (bk BookHelper) BuyBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	count := vars["count"]
	size, _ := strconv.ParseUint(count, 10, 32)
	result, error := bk.Repo.BuyByID(id, uint(size))
	if error != nil {
		http.Error(w, error.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&result)

}

/////

func (bk BookHelper) GetBooksWithPagination(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	offset := vars["offset"]
	pagesize := vars["pagesize"]
	o, _ := strconv.Atoi(offset)
	p, _ := strconv.Atoi(pagesize)
	result, error := bk.Repo.Pagination(o, p)
	if error != nil {
		http.Error(w, error.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&result)

}

////

func (bk BookHelper) GetBooksWithAuthors(w http.ResponseWriter, r *http.Request) {

	result, error := bk.Repo.GetBooksWithAuthor()
	if error != nil {
		http.Error(w, error.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&result)

}

package server

import (
	"encoding/json"
	"net/http"

	. "github.com/hkaya15/PicusSecurity/Week_5_Homework/api/models/api"
	. "github.com/hkaya15/PicusSecurity/Week_5_Homework/app/domain/repository"
)

type AuthorHelper struct {
	Repo *AuthorRepository
}

func NewAuthorHandler(a *AuthorRepository) AuthorHelper {
	return AuthorHelper{Repo: a}
}

func (a AuthorHelper) UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	var book Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	x := APIResponseBook{
		Code:    http.StatusOK,
		Message: "Book is added with success",
		Result:  &book,
	}
	if err := json.NewEncoder(w).Encode(&x); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}



func (a AuthorHelper) FindAuthorById (w http.ResponseWriter, r *http.Request) {
	var book Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	x:=APIResponseBook{
		Code:    http.StatusOK,
		Message: "Book is added with success",
		Result:  &book,
	}
	if err := json.NewEncoder(w).Encode(&x); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
}



func (a AuthorHelper) FindAuthorsByName (w http.ResponseWriter, r *http.Request) {
	var book Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	x:=APIResponseBook{
		Code:    http.StatusOK,
		Message: "Book is added with success",
		Result:  &book,
	}
	if err := json.NewEncoder(w).Encode(&x); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
}




func (a AuthorHelper) GetAllAuthors(w http.ResponseWriter, r *http.Request) {
	var book Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	x:=APIResponseBook{
		Code:    http.StatusOK,
		Message: "Book is added with success",
		Result:  &book,
	}
	if err := json.NewEncoder(w).Encode(&x); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
}

////


func (a AuthorHelper) GetAuthorsWithBooks (w http.ResponseWriter, r *http.Request) {
	var book Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	x:=APIResponseBook{
		Code:    http.StatusOK,
		Message: "Book is added with success",
		Result:  &book,
	}
	if err := json.NewEncoder(w).Encode(&x); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
}
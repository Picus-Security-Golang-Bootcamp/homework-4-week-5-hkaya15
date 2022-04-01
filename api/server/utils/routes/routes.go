package server

import (
	"net/http"

	"github.com/gorilla/mux"
	. "github.com/hkaya15/PicusSecurity/Week_5_Homework/api/server/handlers"
)

var m *mux.Router

func Routes(router *mux.Router, au AuthorHelper, bk BookHelper) {
	router.HandleFunc("/book", bk.CreateBook).Methods(http.MethodPost)
	router.HandleFunc("/book", bk.UpdateBook).Methods(http.MethodPut)
	router.HandleFunc("/book/{id}", bk.FindBookById).Methods(http.MethodGet)
	router.HandleFunc("/book/{id}", bk.DeleteBookById).Methods(http.MethodDelete)
	router.HandleFunc("/bookname/{name}", bk.FindBookByName).Methods(http.MethodGet)
	router.HandleFunc("/book/buy/{id}/count/{count}", bk.BuyBookById).Methods(http.MethodPost)
	router.HandleFunc("/bookswithauthor", bk.GetBooksWithAuthors).Methods(http.MethodGet)
	router.HandleFunc("/books/{offset}/size/{pagesize}", bk.GetBooksWithPagination).Methods(http.MethodGet)
	
	router.HandleFunc("/author", au.UpdateAuthor).Methods(http.MethodPut)
	router.HandleFunc("/author/{authorId}", au.FindAuthorById).Methods(http.MethodGet)
	router.HandleFunc("/author/param?={authorname}", au.FindAuthorsByName).Methods(http.MethodGet)
	router.HandleFunc("/authors", au.GetAllAuthors).Methods(http.MethodGet)
	router.HandleFunc("/authorswithbooks", au.GetAuthorsWithBooks).Methods(http.MethodGet)
}

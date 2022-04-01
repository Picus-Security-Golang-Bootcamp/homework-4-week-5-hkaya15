package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/gorilla/handlers"
	. "github.com/hkaya15/PicusSecurity/Week_5_Homework/api/server/handlers"
	. "github.com/hkaya15/PicusSecurity/Week_5_Homework/api/server/utils/routes"
	. "github.com/hkaya15/PicusSecurity/Week_5_Homework/app/domain/repository"
	. "github.com/hkaya15/PicusSecurity/Week_5_Homework/base/db"
	. "github.com/hkaya15/PicusSecurity/Week_5_Homework/base/logs"
	"github.com/joho/godotenv"
)

func main() {
	logger := CreateLogs()
	defer logger.File.Close()
	err := godotenv.Load()
	if err != nil {
		logger.Logger.Fatalln("Error loading .env file")
	}
	base := DBBase{DbType: &POSTGRES{}}
	db, err := base.DbType.Create()

	if err != nil {
		logger.Logger.Fatalln("DB cannot init")
	}
	logger.Logger.Println("DB connected: ", db)

	bookRepo := NewBookRepository(db)
	authorRepo := NewAuthorRepository(db)

	authorHelper := NewAuthorHandler(authorRepo)
	bookHelper := NewBookHandler(bookRepo)

	authorRepo.Migrations()
	authorRepo.InsertData()
	bookRepo.Migrations()
	bookRepo.InsertData()

	router := mux.NewRouter()

	handlers.AllowedOrigins([]string{"http://127.0.0.1:8080"})
	handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})
	handlers.AllowedMethods([]string{"POST", "GET", "PUT", "DELETE"})

	Routes(router, authorHelper, bookHelper)

	logger.Logger.Println("API is running")
	http.ListenAndServe(":8080", router)
}

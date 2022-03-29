package main

import (
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
	authorRepo.Migrations()
	authorRepo.InsertData()
	bookRepo.Migrations()
	bookRepo.InsertData()
}

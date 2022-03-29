package main

import (
	"log"

	"github.com/joho/godotenv"
	. "github.com/hkaya15/PicusSecurity/Week_5_Homework/base/db"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file")
	}
	base := DBBase{DbType: &POSTGRES{}}
	db, err := base.DbType.Create()

	if err != nil {
		log.Fatalln("DB cannot init")
	}
	log.Println("DB connected: ", db)
}

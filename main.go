package main

import (
	. "github.com/hkaya15/PicusSecurity/Week_5_Homework/base/db"
	. "github.com/hkaya15/PicusSecurity/Week_5_Homework/base/logs"
	"github.com/joho/godotenv"
)

func main() {
	f, logger, err := CreateLogs()
	if err != nil {
		logger.Println(err)
	}
	defer f.Close()
	err = godotenv.Load()
	if err != nil {
		logger.Fatalln("Error loading .env file")
	}
	base := DBBase{DbType: &POSTGRES{}}
	db, err := base.DbType.Create()

	if err != nil {
		logger.Fatalln("DB cannot init")
	}
	logger.Println("DB connected: ", db)
}

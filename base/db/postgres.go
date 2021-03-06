package db

import (
	"fmt"
	"os"
	"sync"

	. "github.com/hkaya15/PicusSecurity/Week_5_Homework/base/logs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type POSTGRES struct{}

var (
	once  sync.Once
	conn *gorm.DB
	log  *BuiltinLogger
)

func (ps *POSTGRES) Create() (*gorm.DB, error) {
	once.Do(func() {
		res, err := CreatePostgreSQL()
		conn=res
		if err != nil {
			log.Logger.Fatalln(err.Error())
		}
	})
	return conn, nil
}

func CreatePostgreSQL() (*gorm.DB, error) {
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		os.Getenv("PICUS_DB_HOST"),
		os.Getenv("PICUS_DB_PORT"),
		os.Getenv("PICUS_DB_USERNAME"),
		os.Getenv("PICUS_DB_NAME"),
		os.Getenv("PICUS_DB_PASSWORD"),
	)
	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		log.Logger.Println(err.Error())
		return nil, fmt.Errorf("cannot open database: %v", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Logger.Println(err.Error())
		return nil, err
	}
	if err := sqlDB.Ping(); err != nil {
		log.Logger.Println(err.Error())
		return nil, err
	}
	return db, nil
}

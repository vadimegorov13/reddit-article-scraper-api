package db

import (
	"log"

	"github.com/vadimegorov13/reddit-article-scraper-api/pkg/common/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("posts.db"), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	} else {
		// this will be printed in the terminal, confirming the connection to the database
		log.Println("The database is connected")
	}

	db.AutoMigrate(&models.Post{})
	log.Println("Dtabase migrated!")

	return db
}

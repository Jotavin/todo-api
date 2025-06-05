package models

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


type Task struct{
	ID int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Done bool `json:"done"`
}

var TaskList = []Task{}


func ConnectDB() *gorm.DB {
	var dsn string = "host=localhost user=admin password=myadmin dbname=todo-api port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	
	if err != nil {
		fmt.Println(err)
        panic(fmt.Sprintf("Failed to connect to database: %v", err))
	}
	return db
}

func MigrateDB(db *gorm.DB) {
	err := db.AutoMigrate(&Task{})
	if err != nil{
		panic(fmt.Sprintf("Failed to migrate database: %v", err))
	}
}
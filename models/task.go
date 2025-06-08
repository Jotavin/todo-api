package models

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


type Task struct{
	ID int64 `json:"id" gorm:"primaryKey"`
	Title string `json:"title"`
	Description string `json:"description"`
	Done bool `json:"done"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

var TaskList = []Task{}


func ConnectDB() *gorm.DB {
	var dsn string = "host=localhost user=admin password=myadmin dbname=todo-api port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
        panic(fmt.Sprintf("Failed to connect to database: %v", err))
	}

	return db
}

func MigrateDB() (string, error){
	db := ConnectDB()
	err := db.AutoMigrate(&Task{})
	if err != nil{
		return "Something occurred in the Migration", fmt.Errorf("failed to migrate database %v", err)
	}

	return "The Migration runned successfully", nil
}
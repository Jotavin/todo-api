package models

import (
	"fmt"
	"os"
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

func ConnectDB() *gorm.DB {
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "localhost" // fallback para desenvolvimento local
	}
	
	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "5432"
	}
	
	user := os.Getenv("DB_USER")
	if user == "" {
		user = "admin"
	}
	
	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		password = "password123"
	}
	
	dbname := os.Getenv("DB_NAME")
	if dbname == "" {
		dbname = "todo-api"
	}
	var dsn string = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)
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
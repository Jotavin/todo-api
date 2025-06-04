package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)
type TaskTest struct{
	gorm.Model
	ID int
	Title string
	Description string
	Done bool
}

func main(){
	CreateDB()
}

func CreateDB(){
	db, err := gorm.Open(sqlite.Open("test.sqlite"), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	db.AutoMigrate(&TaskTest{})
}

func CreateTask(){
	db, err := gorm.Open(sqlite.Open("test.sqlite"), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	db.Create(&TaskTest{ID: 1, Title: "Academia", Description: "Treinar peito", Done: false})
}

func ReadTask(){
	db, err := gorm.Open(sqlite.Open("test.sqlite"), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	var task TaskTest
	db.First(&task, 1)
	db.First(&task, "id = ?", "1")
}

func UpdateTask(){
	db, err := gorm.Open(sqlite.Open("test.sqlite"), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	
	var task TaskTest
	db.Model(&task).Update("Description", "Treinar perna")
	db.Model(&task).Updates(TaskTest{Title: "Workout",Description: "Treinar costas"})
	db.Model(&task).Updates(map[string]interface{}{"Title": "Academia", "Done": true})

}

func DeleteTask() {
	db, err := gorm.Open(sqlite.Open("test.sqlite"), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	
	var task TaskTest
	db.Delete(&task, 1)
}
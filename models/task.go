package models

import (
	"fmt"
	"os"
	"time"
	"todo-api/metrics"

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

var DB *gorm.DB

func getEnvOrDefault(key, defaultValue string){

}

func buildDSN() string{


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
		password = "myadmin"
	}
	
	dbname := os.Getenv("DB_NAME")
	if dbname == "" {
		dbname = "todo-api"
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)
	return dsn
}

func ConnectDB() (*gorm.DB, error){
	if DB != nil {
		return DB, nil
	}

	start := time.Now()
	defer func() {
		metrics.DatabaseConnectionDuration.Observe(time.Since(start).Seconds())
	}()

	metrics.DatabaseConnectionAttempts.Inc()

	dsn := buildDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		metrics.DatabaseConnectionFailures.Inc()
		metrics.DatabaseErrors.WithLabelValues("connect", "connection_failed").Inc()
		
		fmt.Printf("Erro de conexão: %v\n", err)
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	metrics.DatabaseConnectionsActive.Inc()

	return db, nil
}

func MigrateDB() (string, error){
	db, err := ConnectDB()
	if err != nil{
		metrics.RequestsTotal.WithLabelValues("/createTask", "POST", "500").Inc()
		return "", fmt.Errorf("failed to connect to database: %v", err)
	}

	err = db.AutoMigrate(&Task{})
	if err != nil{
		return "Something occurred in the Migration", fmt.Errorf("failed to migrate database %v", err)
	}

	return "The Migration runned successfully", nil
}
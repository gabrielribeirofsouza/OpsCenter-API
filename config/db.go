package config

import (
	"database/sql"
	"os"
	"fmt"
	"github.com/joho/godotenv"
)

func SetupDB() (*sql.DB, error){
	err := godotenv.Load()
	if err != nil{
		return nil, err
	}
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")	
	dbHost := os.Getenv("DB_HOST")	
	dbPort := os.Getenv("DB_PORT")	
	dbSSLmode := os.Getenv("DB_SSLMODE")

	connectionStr := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s", dbHost, dbPort, dbName, dbUser, dbPassword, dbSSLmode)
	dbConnection, err := sql.Open("postgres", connectionStr)
	if err != nil{
		return nil, err
	}
	err = dbConnection.Ping()
	if err != nil{
		return nil, err
	}
	return dbConnection, nil
}
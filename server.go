package main

import (
	"JobScheduler/Server/database"
	"JobScheduler/Server/model"
	"log"
	"os"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	//	r := mux.NewRouter()

	// fmt.Printf("Starting server at port 8080\n")

	// srv := &http.Server{
	// 	Addr:    ":8080",
	// 	Handler: r,
	// }
	// srv.ListenAndServe()
	initDB()
}

func initDB() {
	config := database.Config{
		Host:     os.Getenv("HOST"),
		User:     os.Getenv("USER"),
		Password: os.Getenv("PASSWORD"),
		DB:       os.Getenv("DB"),
	}

	err := database.Connect(config)
	if err != nil {
		panic(err.Error())
	}

	database.Migrate(&model.Job{})
}

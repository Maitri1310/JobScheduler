package main

import (
	"JobScheduler/Server/controller"
	"JobScheduler/Server/database"
	"JobScheduler/Server/model"
	"JobScheduler/Server/scheduler"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

func main() {
	defer database.Session.Close()
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	initDB()

	router := mux.NewRouter().StrictSlash(true)
	initRoute(router)
	database.InitialiseRedisWorker()
	database.CreateConnection()
	go scheduler.Scheduler()
	fmt.Printf("Starting server at port 8080\n")
	http.ListenAndServe(":8080", router)

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

func initRoute(router *mux.Router) {
	router.HandleFunc("/jobs/{id}", controller.FindJob).Methods("GET")
	router.HandleFunc("/jobs", controller.CreateJob).Methods("POST")
	router.HandleFunc("/jobs/{id}", controller.DeleteJob).Methods("DELETE")
}

package main

import (
	"JobScheduler/Server/database"
	"JobScheduler/Server/model"
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	//	r := mux.NewRouter()

	fmt.Printf("Starting server at port 8080\n")

	// srv := &http.Server{
	// 	Addr:    ":8080",
	// 	Handler: r,
	// }
	// srv.ListenAndServe()
	initDB()
}

func initDB() {

	config := database.Config{
		Host:     "localhost",
		User:     "root",
		Password: "admin",
		DB:       "job_scheduler",
	}

	err := database.Connect(config)
	if err != nil {
		panic(err.Error())
	}

	database.Migrate(&model.Job{})

}

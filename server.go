package main

import (
	"JobScheduler/Server/controller"
	"JobScheduler/Server/database"
	"JobScheduler/Server/model"
	"JobScheduler/Server/scheduler"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

func closeServer(wg sync.WaitGroup, ctx context.Context, server *http.Server) {
	defer wg.Done()
	<-ctx.Done()
	log.Println("Closing HTTP Server")
	if err := server.Shutdown(context.Background()); err != nil {
		log.Fatal(err)
	}
}
func createServer(wg sync.WaitGroup, server *http.Server) {
	defer wg.Done()
	fmt.Printf("Starting server at port 8080\n")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Server started")
}
func main() {
	defer database.Session.Close()

	err := godotenv.Load(".env")
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	var wg sync.WaitGroup
	wg.Add(4)
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	initMysqlDB()

	database.CreateCassandraConnection()
	go scheduler.Scheduler(wg, ctx)
	go database.InitialiseRedisWorker(wg, ctx)
	router := mux.NewRouter().StrictSlash(true)

	initRoute(router)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	go createServer(wg, server)

	go closeServer(wg, ctx, server)
	wg.Wait()
}

func initMysqlDB() {
	config := database.Config{
		Host:     os.Getenv("HOST"),
		User:     os.Getenv("MYSQLUSER"),
		Password: os.Getenv("PASSWORD"),
		DB:       os.Getenv("DB"),
	}

	err := database.Connect(config)
	if err != nil {
		log.Fatal(err)
		panic(err.Error())
	}

	database.Migrate(&model.Job{})
}

func initRoute(router *mux.Router) {
	router.HandleFunc("/jobs/{id}", controller.FindJob).Methods("GET")
	router.HandleFunc("/jobs", controller.CreateJob).Methods("POST")
	router.HandleFunc("/jobs/{id}", controller.DeleteJob).Methods("DELETE")
}

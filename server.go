package main

import (
	"JobScheduler/Server/middleware"
	"JobScheduler/Server/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/jobs", handlePost).Methods("POST")
	r.HandleFunc("/jobs/{id}", handleGet).Methods("GET")

	fmt.Printf("Starting server at port 8080\n")
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	srv.ListenAndServe()
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// create an empty user of type models.User
	var user models.Job

	// decode the json request to user
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	// call insert user function and pass the user
	insertID := middleware.Insert(user)

	// format a response object
	fmt.Printf(insertID)

	res := response{}

	// send the response
	json.NewEncoder(w).Encode(res)
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// get the userid from the request params, key is "id"
	params := mux.Vars(r)

	// convert the id type from string to int
	id := params["id"]

	// call the getUser function with user id to retrieve a single user
	job, err := middleware.Retrive(id)

	if err != nil {
		log.Fatalf("Unable to get user. %v", err)
	}

	// send the response
	json.NewEncoder(w).Encode(job)
}

type response struct {
	ID int64 `json:"id,omitempty"`
}

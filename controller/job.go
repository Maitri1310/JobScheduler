package controller

import (
	"JobScheduler/Server/database"
	"JobScheduler/Server/model"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func CreateJob(w http.ResponseWriter, r *http.Request) {

	requestBody, _ := ioutil.ReadAll(r.Body)

	var job model.Job

	json.Unmarshal(requestBody, &job)
	job.Id = uuid.New().String()

	database.Connector.Create(job)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(job)

}
func DeleteJob(w http.ResponseWriter, r *http.Request) {
	m := mux.Vars(r)

	id := m["id"]
	var job model.Job

	database.Connector.Where("Id=?", id).Delete(&job)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(job)

}
func FindJob(w http.ResponseWriter, r *http.Request) {
	m := mux.Vars(r)

	id := m["id"]
	var job model.Job
	database.Connector.First(&job, id)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(job)

}

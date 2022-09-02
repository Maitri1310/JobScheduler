package controller

import (
	"JobScheduler/Server/database"
	"JobScheduler/Server/model"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type deleteResponse struct {
	Id string
}

func CreateJob(w http.ResponseWriter, r *http.Request) {

	requestBody, _ := ioutil.ReadAll(r.Body)

	var job model.Job

	json.Unmarshal(requestBody, &job)
	job.Id = uuid.New().String()

	database.Connector.Create(job)
	now := time.Now()

	database.Session.Query(`INSERT INTO job_scheduler.jobPool (interval, nextRunTime, jobId, webhook) VALUES (?, ?, ?, ?)`, job.Cron, now.Unix(), job.Id, job.Webhook).Exec()

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(job)

}
func DeleteJob(w http.ResponseWriter, r *http.Request) {
	m := mux.Vars(r)

	id := m["id"]
	var job = model.Job{Id: id}

	result := database.Connector.Delete(&job)

	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if result.RowsAffected != 1 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	database.Session.Query(`DELETE FROM job_scheduler.jobPool WHERE jobId = ?`, id).Exec()
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(deleteResponse{Id: id})

}
func FindJob(w http.ResponseWriter, r *http.Request) {
	m := mux.Vars(r)
	id := m["id"]

	var job = model.Job{Id: id}

	database.Connector.First(&job)

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(job)

}

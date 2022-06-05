package middleware

import (
	"JobScheduler/Server/models"
	"fmt"

	"github.com/google/uuid"
)

var ir_MAP = map[string]models.Job{}

var ErrIDNotFound = fmt.Errorf("Id not found")

func Insert(job models.Job) string {
	job.Id = uuid.New().String()
	ir_MAP[job.Id] = job

	fmt.Println(job)
	return job.Id
}

func Retrive(id string) (models.Job, error) {
	v, found := ir_MAP[id]

	if !found {
		return models.Job{}, ErrIDNotFound
	}

	return v, nil
}

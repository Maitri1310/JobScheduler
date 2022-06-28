package scheduler

import (
	"JobScheduler/Server/database"
	"log"
	"time"

	"github.com/gocraft/work"
)

var enqueuer = work.NewEnqueuer("job_scheduler", database.RedisPool)

func EnqueueJob(webhook string) {
	_, err := enqueuer.Enqueue("take_job", work.Q{"webhook": webhook})
	if err != nil {
		log.Fatal(err)
	}
}
func Scheduler() {

	for {
		time.Sleep(60 * time.Second)
	}
    
}

func GetJobs(){

	
}

package scheduler

import (
	"JobScheduler/Server/database"
	"fmt"
	"log"

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
	fmt.Println("Starting scheduler")
	//for {
	//time.Sleep(5 * time.Second)
	fmt.Println("5Sec passed")
	go GetJobs()
	//}

}

func IterateOverJobs() {

}

func GetJobs() {
	var webhook string
	var nextRunTime int32
	var interval int32
	var jobId string

	iter := database.Session.Query(`SELECT webhook, nextRunTime, interval, jobId FROM jobPool WHERE nextRunTime <= toTimestamp(now()) ALLOW FILTERING`).Iter()
	for iter.Scan(&webhook, &nextRunTime, &interval, &jobId) {
		fmt.Println("Tweet:", webhook, nextRunTime, interval, jobId)
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}
}

func UpdateJob() {

}

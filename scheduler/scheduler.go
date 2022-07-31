package scheduler

import (
	"JobScheduler/Server/database"
	"fmt"
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
	fmt.Println("Starting scheduler")
	// signalChan := make(chan os.Signal, 1)
	// signal.Notify(signalChan, os.Interrupt, os.Kill)

	// ticker := time.NewTicker(5 * time.Second)
	// loop := true
	// for loop {
	// 	select {
	// 	// Got a timeout! fail with a timeout error
	// 	case <-signalChan:
	// 		fmt.Println("finish")
	// 		loop = false
	// 		break
	// 	// Got a tick, we should check on checkSomething()
	// 	case <-ticker.C:
	// 		fmt.Println("5 sec passed")
	// 		go GetJobs()
	// 		// checkSomething() isn't done yet, but it didn't fail either, let's try again
	// 	}
	// }
	for tick := range time.Tick(5 * time.Second) {
		fmt.Println("60 sec passed", tick)
		go GetJobs()
	}
	// close(signalChan)
	// ticker.Stop()
	// fmt.Println("exit for")
	// return
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
		go UpdateJob(jobId, nextRunTime, interval)
		go EnqueueJob(webhook)
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}
}

func UpdateJob(jobid string, nextRuntime int32, interval int32) {

	err := database.Session.Query(`UPDATE job_scheduler.jobPool SET nextRuntime=?+? where jobId=?`, nextRuntime, interval, jobid).Exec()
	if err != nil {
		log.Fatal(err)
	}

}

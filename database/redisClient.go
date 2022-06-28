package database

import (
	"JobScheduler/Server/worker"
	"fmt"
	"os"
	"os/signal"

	"github.com/gocraft/work"
	"github.com/gomodule/redigo/redis"
)

// Make a redis pool
var RedisPool = &redis.Pool{
	MaxActive: 5,
	MaxIdle:   5,
	Wait:      true,
	Dial: func() (redis.Conn, error) {
		return redis.Dial("tcp", ":6379")
	},
}

type Context struct {
	customerID int64
}

func InitialiseRedisWorker() {
	pool := work.NewWorkerPool(Context{}, 10, "job_scheduler", RedisPool)

	// Add middleware that will be executed for each job
	pool.Middleware((*Context).Log)

	// Map the name of jobs to handler functions
	pool.Job("take_job", (*Context).TakeJob)

	// Customize options:
	pool.JobWithOptions("export", work.JobOptions{Priority: 10, MaxFails: 1}, (*Context).Export)

	// Start processing jobs
	pool.Start()

	// Wait for a signal to quit:
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, os.Kill)
	<-signalChan

	// Stop the pool
	pool.Stop()
}
func (c *Context) Log(job *work.Job, next work.NextMiddlewareFunc) error {
	fmt.Println("Starting job: ", job.Name)
	return next()
}
func (c *Context) TakeJob(job *work.Job) error {
	webhook := job.ArgString("webhook")
	if err := job.ArgError(); err != nil {
		return err
	}

	go worker.ProceessJob(webhook)

	return nil
}
func (c *Context) Export(job *work.Job) error {
	return nil
}

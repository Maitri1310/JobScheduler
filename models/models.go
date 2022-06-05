package models

type Job struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Webhook string `json:"webhook"`
	Cron    string `json:"cron"`
}

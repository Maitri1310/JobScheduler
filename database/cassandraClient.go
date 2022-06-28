package database

import (
	"log"

	"github.com/gocql/gocql"
)

var Session *gocql.Session

func CreateConnection() {
	cluster := gocql.NewCluster("127.0.0.1:9042")
	cluster.Keyspace = "job_scheduler"
	cluster.Consistency = gocql.Quorum
	var err error
	Session, err = cluster.CreateSession()
	if err != nil {
		log.Fatal(err)
	}
}

package database

import (
	"log"
	"time"

	"github.com/gocql/gocql"
)

var Session *gocql.Session

func CreateCassandraConnection() {
	cluster := gocql.NewCluster("127.0.0.1:9042")
	cluster.Keyspace = "job_scheduler"
	cluster.Consistency = gocql.Quorum
	cluster.ConnectTimeout = time.Second * 10
	var err error
	Session, err = cluster.CreateSession()
	if err != nil {
		log.Fatal(err)
	}
}

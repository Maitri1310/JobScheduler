package worker

import (
	"io/ioutil"
	"log"
	"net/http"
)

type Context struct {
	customerID int64
}

func ProceessJob(webhook string) {
	res, err := http.Get(webhook)
	if err != nil {
		log.Println(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}
	//Convert the body to type string
	sb := string(body)
	log.Printf(sb)

}

// func (c *Context) TakeJob(job *work.Job) error {
// 	webhook := job.ArgString("webhook")
// 	if err := job.ArgError(); err != nil {
// 		return err
// 	}

// 	go ProceessJob(webhook)

// 	return nil
// }

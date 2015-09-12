package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type HttpResp struct {
	name   string
	url    string
	body   string
	status string
	err    error
}

type Page struct {
	Services []*HttpResp //map[string]map[string]string
}

var services = map[string]map[string]string{
	"micropig": {
		"url": "http://localhost:12312/micropig", //"http://streaker.technoblogic.io/micropig",
	},
	"microscope": {
		"url": "http://localhost:12312/microscope", //"http://streaker.technoblogic.io/microscope",
	},
	"microbrew": {
		"url": "http://localhost:12312/microbrew", //"http://streaker.technoblogic.io/microbrew",
	},
}

func asyncQuery(services map[string]map[string]string) []*HttpResp {
	// A channel for responses
	respCh := make(chan *HttpResp)
	// Array of responses
	responses := []*HttpResp{}
	// A loop with a nested go func to feed the channel with the results of the query
	for service, _ := range services {
		url := services[service]["url"]
		go func(url string) {
			log.Info("Fetching: ", url)
			resp, err := http.Get(url)
			if err != nil {
				log.Warn(err)
			}
			body, err := ioutil.ReadAll(resp.Body)
			status := resp.Status
			respCh <- &HttpResp{service, url, string(body), status, err}
		}(url)
	}
	// Fill an array with the responses
	for {
		select {
		case r := <-respCh:
			log.Info("Fetched: ", r.name, " service: ", r.body, " - ", r.status)
			// In order to properly break loop, count the number of responses by adding to array
			responses = append(responses, r)
			if len(responses) == 3 {
				return responses
			}
		case <-time.After(time.Millisecond * 50):
			fmt.Printf(".")
		}
	}
	return []*HttpResp{}
}

func Streaker(w http.ResponseWriter, req *http.Request) {
	svcData := asyncQuery(services)
	// Add the data to the page struct for use in template
	var p = &Page{
		Services: svcData,
	}
	// Execute the template
	t, _ := template.ParseFiles("microservices.html")
	t.Execute(w, &p)
}

func main() {
	log.Info("STREAKERs ON :8000!")
	//Handle request
	http.HandleFunc("/", Streaker)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Error("Can't start streaker!")
		log.Error(err)
		os.Exit(1)
	}
}

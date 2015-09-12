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
	name     string
	url      string
	response *http.Response
	err      error
}

type Page struct {
	Services map[string]map[string]string
}

var services = map[string]map[string]string{
	"micropig": {
		"url":    "http://localhost:12312/micropig", //"http://streaker.technoblogic.io/micropig",
		"resp":   "",
		"status": "",
	},
	"microscope": {
		"url":    "http://localhost:12312/microscope", //"http://streaker.technoblogic.io/microscope",
		"resp":   "",
		"status": "",
	},
	"microbrew": {
		"url":    "http://localhost:12312/microbrew", //"http://streaker.technoblogic.io/microbrew",
		"resp":   "",
		"status": "",
	},
}

func asyncQuery(services map[string]map[string]string) map[string]map[string]string {
	// A channel for responses
	respCh := make(chan *HttpResp)
	// Array of responses
	responses := []string{}
	// A loop with a nested go func to feed the channel with the results of the query
	for service, _ := range services {
		url := services[service]["url"]
		go func(url string) {
			log.Info("Fetching: ", url)
			resp, err := http.Get(url)
			respCh <- &HttpResp{service, url, resp, err}
		}(url)
	}
	// Fill an array with the responses
	for {
		select {
		case r := <-respCh:
			respString, err := ioutil.ReadAll(r.response.Body)
			log.Info("Fetched: ", r.url, " for ", r.name, " service: ", string(respString), " - ", r.response.Status)
			if err != nil {
				log.Error(err)
				os.Exit(1)
			}
			services[r.name]["resp"] = string(respString)
			services[r.name]["status"] = r.response.Status
			// In order to properly break loop, count the number of responses by adding to array
			responses = append(responses, r.name)
			if len(responses) == len(services) {
				return services
			}
		case <-time.After(time.Millisecond * 500):
			fmt.Printf(".")
		}
	}
	return services
}

func Streaker(w http.ResponseWriter, req *http.Request) {
	svcData := asyncQuery(services)
	log.Info("Received data:")
	for service, data := range svcData {
		log.Info(service, ": ", data["resp"], " ", data["status"])
	}

	var p = &Page{
		Services: svcData,
	}

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

// +build linux darwin
package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
	"text/template"
	"time"
)

type HttpResp struct {
	Name   string
	Url    string
	Resp   string
	Status string
	Err    error
}

type Page struct {
	Services []*HttpResp //map[string]map[string]string
}

var services = map[string]map[string]string{
	"micropig": {
		"url": "http://localhost:12313/micropig", //"http://streaker.technoblogic.io/micropig",
	},
	"microscope": {
		"url": "http://localhost:12314/microscope", //"http://streaker.technoblogic.io/microscope",
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
		name := service
		go func(url string) {
			log.Info("Fetching: ", url)
			resp, err := http.Get(url)
			if resp == nil {
				body := "Down"
				var dump = &HttpResp{
					Name:   name,
					Url:    url,
					Resp:   body,
					Status: body,
					Err:    err,
				}
				log.Warn("Dumping ", dump.Name, " as ", dump.Resp)
				respCh <- dump

			} else {
				body, _ := ioutil.ReadAll(resp.Body)
				status := resp.Status
				var dump = &HttpResp{
					Name:   name,
					Url:    url,
					Resp:   string(body),
					Status: status,
					Err:    err,
				}
				log.Warn("Dumping ", dump.Name, " as ", dump.Resp)
				respCh <- dump
			}
		}(url)
	}
	// Fill an array with the responses
	for {
		select {
		case r := <-respCh:
			log.Info("Fetched: ", r.Name, " service: ", r.Resp, " - ", r.Status)
			// In order to properly break loop, count the number of responses by adding to array
			responses = append(responses, r)
			if len(responses) == len(services) {
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

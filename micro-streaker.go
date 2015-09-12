package main

import (
	log "github.com/Sirupsen/logrus"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

type HttpResp struct {
	name     string
	url      string
	response *http.Response
	err      error
}

var services = map[string]map[string]string{
	"micropig": {
		"url":  "http://localhost:12312/microbrew", //"http://streaker.technoblogic.io/micropig",
		"resp": "",
	},
	"microscope": {
		"url":  "http://localhost:12312/microscope", //"http://streaker.technoblogic.io/microscope",
		"resp": "",
	},
	"microbrew": {
		"url":  "http://localhost:12312/microbrew", //"http://streaker.technoblogic.io/microbrew",
		"resp": "",
	},
}

func asyncQuery(services map[string]map[string]string) map[string]map[string]string {
	// A channel for responses
	respCh := make(chan *HttpResp)
	// The struct to handle the response
	responses := []*HttpResp{}
	// A loop with a nested go func to feed the channel with the results of the query
	for service, data := range services {
		url := service["url"]
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
			log.Info("Fetched: ", r.url, " for ", r.name, " service.")
			services[r.name]["resp"] = r.resp
			if len(responses) == len(services) {
				return services
			}
		case <-time.After(time.Milisecond * 50):
			log.Info(".")
		}
	}

}

func Streaker(w http.ResponseWriter, req *http.Request) {
	svcData := asyncQuery
	log.Info(svcData)
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, p)
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

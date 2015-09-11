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
	url      string
	response *http.Response
	err      error
}

var SvcUrls = []string{
	"http://streaker.technoblogic.io/micropig",
	"http://streaker.technoblogic.io/microscope",
	"http://streaker.technoblogic.io/microbrew",
}

func asyncQuery(urls []string) []*HttpResp {
	// A channel for responses
	respCh := make(chan *HttpResp)
	// The struct to handle the response
	responses := []*HttpResp{}
	// A loop with a nested go func to feed the channel with the results of the query
	for _, url := range urls {
		go func(url string) {
			log.Info("Fetching: ", url)
			resp, err := http.Get(url)
			ch <- &HttpResp{url, resp, err}
		}(url)
	}
	// For each result in the channel, return a resp
	select {
	case r := <-ch:
		log.Info("Fetched: ", r.url)
		responses = append(responses, r)
		if len(responses) == len(urls) {
			return responses
		}
	}
}

func Streaker(w http.ResponseWriter, req *http.Request) {
	p, err := getData
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

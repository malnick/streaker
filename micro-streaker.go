package main

import (
	log "github.com/Sirupsen/logrus"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

type SvcData struct {
	URL  string
	Resp map[string]string
}

type Services struct {
	Micropig   SvcData
	Microscope SvcData
	Microbrew  SvcData
}

type SvcUrls struct {
	Micropig   "http://streaker.technoblogic.io/micropig"
	Microscope "http://streaker.technoblogic.io/microscope"
	Microbrew  "http://streaker.technoblogic.io/microbrew"
}

func getData() (s Services, err error) {
	// Create channel for resp

	// Go routine to query URL

	return &s
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

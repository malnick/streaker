package main

import (
	log "github.com/Sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func Streaker(w http.ResponseWriter, req *http.Request) {
	index, _ := ioutil.ReadFile("./index.html")
	io.WriteString(w, string(index))
}

func main() {
	//Handle request
	http.HandleFunc("/", Streaker)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Error("Can't start streaker!")
		log.Error(err)
		os.Exit(1)
	}
}

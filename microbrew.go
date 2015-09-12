package main

import (
	"bufio"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func response(rw http.ResponseWriter, request *http.Request) {

	if err != nil {
		log.Println("ERROR: ", err)
	}
	json, err := json.Marshal(services)
	rw.Write([]byte("http://cdn.psfk.com/wp-content/uploads/2013/07/beer-labels-dogfish-head.gif"))
}

func main() {
	http.HandleFunc("/microbrew", response)
	http.ListenAndServe(":12312", nil)
}

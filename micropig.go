package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"os"
)

func response(rw http.ResponseWriter, request *http.Request) {
	pigGifs := []string{
		`http://pix.avaxnews.com/avaxnews/12/42/00014212_medium.jpeg`,
		`http://ak-hdl.buzzfed.com/static/enhanced/webdr06/2013/7/10/16/enhanced-buzz-2152-1373487997-50.jpg`,
		`http://ak-hdl.buzzfed.com/static/enhanced/webdr03/2013/7/10/16/enhanced-buzz-26573-1373487998-27.jpg`,
		`http://ak-hdl.buzzfed.com/static/enhanced/webdr03/2013/7/10/16/enhanced-buzz-26830-1373488000-23.jpg`,
		`http://ak-hdl.buzzfed.com/static/enhanced/webdr03/2013/7/10/16/enhanced-buzz-26702-1373488003-14.jpg`,
		`http://ak-hdl.buzzfed.com/static/enhanced/webdr05/2013/7/10/16/enhanced-buzz-12350-1373488004-24.jpg`,
		`http://ak-hdl.buzzfed.com/static/enhanced/webdr03/2013/7/10/16/enhanced-buzz-26914-1373487996-17.jpg`,
		`http://ak-hdl.buzzfed.com/static/enhanced/webdr03/2013/7/10/16/enhanced-buzz-26077-1373488008-26.jpg`,
	}

	returnUrl := pigGifs[rand.Intn(len(pigGifs))]
	json, err := json.Marshal(returnUrl)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	rw.Write([]byte(json))
}

func main() {
	log.Println("Micropig squelling on localhost:12312/micropig")
	http.HandleFunc("/micropig", response)
	http.ListenAndServe(":12313", nil)
}

package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

func response(rw http.ResponseWriter, request *http.Request) {
	beerGifs := []string{
		`http://hypefreshmag.com/wp-content/uploads/2015/06/beer-genes-app-will-help-people-pick-beers.jpg`,
		`https://lygsbtd.files.wordpress.com/2011/08/beer_toast.jpg`,
		`http://rule13.com/beerrun/hophead/wp-content/uploads/2014/01/taps.jpg`,
	}

	for _, value := range beerGifs {
		json, err := json.Marshal(value)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		rw.Write([]byte(json))
		time.Sleep(time.Second * 1)
	}
}

func main() {
	log.Println("Microbrew tapped on localhost:12312/microbrew")
	http.HandleFunc("/microbrew", response)
	http.ListenAndServe(":12312", nil)
}

package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"os"
)

func response(rw http.ResponseWriter, request *http.Request) {
	beerGifs := []string{
		`http://vinobmo.com/wp-content/uploads/beer_wallpaper_by_gruberjan.jpg`,
		`http://www.milliokerestaurant.com/assets/images/main/ml_main-beer-pic.jpg`,
		`http://i.huffpost.com/gen/1406344/images/o-BEER-facebook.jpg`,
		`http://www.chicagomag.com/Chicago-Magazine/September-2013/best-beer/chicagomag-beer-desktop-1920x1200.jpg`,
		`https://s3.amazonaws.com/images1.vat19.com/beer-boot/beer-boot-40-ounces.jpg`,
		`http://hypefreshmag.com/wp-content/uploads/2015/06/beer-genes-app-will-help-people-pick-beers.jpg`,
		`https://lygsbtd.files.wordpress.com/2011/08/beer_toast.jpg`,
		`http://rule13.com/beerrun/hophead/wp-content/uploads/2014/01/taps.jpg`,
	}

	returnUrl := beerGifs[rand.Intn(len(beerGifs))]
	json, err := json.Marshal(returnUrl)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	rw.Write([]byte(json))
}

func main() {
	log.Println("Microbrew tapped on localhost:12312/microbrew")
	http.HandleFunc("/microbrew", response)
	http.ListenAndServe(":12312", nil)
}

package main

import (
	//	"encoding/json"
	"log"
	"net/http"
	//	"os"
)

func response(rw http.ResponseWriter, request *http.Request) {
	beerGif := `//giphy.com/embed/zrj0yPfw3kGTS`
	//	json, err := json.Marshal(beerGif)
	//	if err != nil {
	//	log.Println(err)
	//	os.Exit(1)
	//}
	rw.Write([]byte(beerGif))
}

func main() {
	log.Println("Microbrew tapped on localhost:12312/microbrew")
	http.HandleFunc("/microbrew", response)
	http.ListenAndServe(":12312", nil)
}

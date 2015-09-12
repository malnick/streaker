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
		`<div style="max-width: 500px;" id="_giphy_cHw5gruhGb0IM"></div><script>var _giphy = _giphy || []; _giphy.push({id: "cHw5gruhGb0IM",w: 500, h: 375});var g = document.createElement("script"); g.type = "text/javascript"; g.async = true;g.src = ("https:" == document.location.protocol ? "https://" : "http://") + "giphy.com/static/js/widgets/embed.js";var s = document.getElementsByTagName("script")[0]; s.parentNode.insertBefore(g, s);</script>`,
	}

	for _, value := range beerGifs {
		json, err := json.Marshal(value)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		rw.Write([]byte(value))
		time.Sleep(time.Second * 1)
	}
}

func main() {
	log.Println("Microbrew tapped on localhost:12312/microbrew")
	http.HandleFunc("/microbrew", response)
	http.ListenAndServe(":12312", nil)
}

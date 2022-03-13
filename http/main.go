package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", (func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{\"message\": \"Hello World\"}"))
	}))
	http.HandleFunc("/getvideos", HandleGetVideos)

	http.HandleFunc("/updatevideos", HandleUpdateVideos)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func HandleGetVideos(w http.ResponseWriter, r *http.Request) {

	videos := getVideos()

	videoBytes, err := json.Marshal(videos)

	if err != nil {
		panic(err)
	}

	w.Write(videoBytes)

}

func HandleUpdateVideos(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			panic(err)
		}

		var videos []video

		err = json.Unmarshal(body, &videos)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Error: %s", err)
		}

		saveVideos(videos)

	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("{\"message\": \"Method not allowed\"}"))
	}

}

package main

import "fmt"

func main() {
	videos := getVideos()

	for index := range videos {
		videos[index].Description = videos[index].Description + "\n Remeber to Like and Subscribe to my channel"
	}
	fmt.Println(videos)

	saveVideos(videos) 
}

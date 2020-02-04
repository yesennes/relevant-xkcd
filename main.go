package main

import (
	"fmt"
	"net/http"
	"time"
)

var comics []XKCDComic

func main() {
	loadComics()

	go func() {
		for {
			fmt.Printf("Started crawl\n")
			crawl()
			fmt.Printf("Crawl finished\n")
			loadComics()
			time.Sleep(crawlInterval)
		}
	}()
	fmt.Printf("Started background crawl\n")

	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/search", SearchHandler)

	http.ListenAndServe(":7070", nil)
}

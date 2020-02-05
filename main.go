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
			fmt.Println("Sleeping for ", crawlInterval)
			time.Sleep(crawlInterval)
			fmt.Printf("Started crawl\n")
			crawl()
			fmt.Printf("Crawl finished\n")
			loadComics()
		}
	}()
	fmt.Printf("Started background crawl\n")

	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/search", SearchHandler)

	http.ListenAndServe(":7070", nil)
}

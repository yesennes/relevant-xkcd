package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func crawl() {
	var wg sync.WaitGroup
	tmp_comics := make([]XKCDComic, 0)
	mux := &sync.Mutex{}

	for _, URL := range URLS {
		resp, err := http.Get(URL)
		if err != nil {
			fmt.Println(err)
			return
		}
		if resp.StatusCode != 200 {
			fmt.Printf("Request failed with %s\n", resp.Status)
			return
		}
		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			fmt.Println(err)
			return
		}

		doc.Find("tr").Each(func(i int, row *goquery.Selection) {
			wg.Add(1)
			time.Sleep(requestDelay)
			go func() {
				defer wg.Done()

				comic := XKCDComic{}
				explanationURL := ""

				row.Find("td").Each(func(j int, col *goquery.Selection) {
					text := strings.TrimSpace(col.Text())

					switch j {
					case 0:
						comic.URL = text
						comic.Number, _ = strconv.Atoi(text[strings.Index(text, "/")+1:])

					case 1:
						comic.Title = strings.TrimSpace(text[:strings.Index(text, "(create)")-1])
						comic.TitleFields = strings.Fields(comic.Title)

						explanationURL, _ = col.Find("a").Attr("href")
						explanationURL = "http://www.explainxkcd.com" + explanationURL[:15] + "?action=edit&title=" + explanationURL[16:]

						resp, err := http.Get(explanationURL)
						if err != nil {
							fmt.Printf("Request for comic %s", explanationURL)
							fmt.Println(err)
							return
						}
						if resp.StatusCode != 200 {
							fmt.Printf("Request for comic %s failed with %s\n", explanationURL, resp.Status)
							return
						}
						exp, err := goquery.NewDocumentFromReader(resp.Body)
						if err == nil {
							comic.Text = exp.Find("textarea").Text()
						} else {
							fmt.Println(err)
						}

					case 3:
						comic.Image = "https://imgs.xkcd.com/comics/" + strings.Replace(text, " ", "_", -1)

					case 4:
						comic.Date = text
					}
				})

				index := strings.Index(comic.Text, "titletext = ")
				if index > 0 {
					comic.TitleText = comic.Text[index+12:]
					comic.TitleText = comic.TitleText[:strings.Index(comic.TitleText, "}")-1]
				}

				mux.Lock()
				tmp_comics = append(tmp_comics, comic)
				mux.Unlock()
			}()
		})
		wg.Wait()
	}

	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	enc.Encode(tmp_comics)

	f, _ := os.Create("comics.bin")
	f.Write(buf.Bytes())
	f.Close()

	comics = tmp_comics
}

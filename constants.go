package main

import (
	"time"
)

var TitleWordWeight int = 400
var TitleIntextWeight int = 200
var TranscriptWeight int = 30
var TextWordWeight int = 8
var TextWeight int = 1

var URLS = []string{
	"http://www.explainxkcd.com/wiki/index.php?title=List_of_all_comics_(1-500)&printable=yes",
	"http://www.explainxkcd.com/wiki/index.php?title=List_of_all_comics_(501-1000)&printable=yes",
	"http://www.explainxkcd.com/wiki/index.php?title=List_of_all_comics_(1001-1500)&printable=yes",
	"http://www.explainxkcd.com/wiki/index.php?title=List_of_all_comics_(1501-2000)&printable=yes",
	"http://www.explainxkcd.com/wiki/index.php?title=List_of_all_comics_&printable=yes",
}

var crawlInterval = 12 * time.Hour
var requestDelay = time.Second

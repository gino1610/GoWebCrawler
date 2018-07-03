package main


import (
	"net/http"
	"fmt"
	"log"
	"io/ioutil"
	"os"
  )


type crawler struct {
	currentPageUrl urlRepo
	externalPageUrl urlRepo
	failedUrlRepo	urlRepo
	otherUrl		urlRepo
	pages 	[]Page
	isCurrentPage	bool
}

type PageList []Page
var crawlerMaster crawler

func PageHasBeenCrawled(url string) bool {
	for _, e := range crawlerMaster.pages {
		if (e.url == url) {
			return true;
		}
	}
	return false;
}

func GetWebText(url string) string {
	var result string
	response, err := http.Get(url)
	if err != nil {
			log.Fatal(err)
	} else {
			defer response.Body.Close()
			contents, err := ioutil.ReadAll(response.Body)
			if err != nil {
				fmt.Printf("%s", err)
				os.Exit(1)
			}
			fmt.Printf("%s\n", string(contents))
			result = string(contents)	
	}
	return result
} 


func crawlPage(url string) {
	isPageAlreadyCrawled := PageHasBeenCrawled(url)
	if(!isPageAlreadyCrawled) {
		htmlText := GetWebText(url)
		fmt.Printf("%v\n", htmlText)
	}
}
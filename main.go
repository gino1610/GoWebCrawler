package main

import (
	"fmt"
	"time"
)

func main() {
	topLevelUrl := "https://www.redhat.com"

	date := time.Now()
	dateformatted := fmt.Sprintf("%d%02d%02d_%02d%02d%02d",
		date.Year(), date.Month(), date.Day(),
		date.Hour(), date.Minute(), date.Second())

	reportFile1 := "D:\\Temp\\GoWebCrawlerReport_" + dateformatted + ".txt"
	reportFile2 := "D:\\Temp\\GoWebCrawlerReport2_" + dateformatted + ".txt"

	fmt.Printf("TopLevelUrl: %v.\n", topLevelUrl)
	fmt.Printf("reportfile1: %v.\n", reportFile1)
	fmt.Printf("reportfile2: %v.\n", reportFile2)
	fmt.Printf("dateFormatted: %s.\n", dateformatted)

	Crawl(topLevelUrl, reportFile1, reportFile2)
}

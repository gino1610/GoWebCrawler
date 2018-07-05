package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

var NodeDB Node
var TopLevelUrl string
var CrawledPages []string
var LinkDB []string

func Crawl(topLevelUrl string, reportFile1 string, reportFile2 string) {

	TopLevelUrl = topLevelUrl
	robotsFile := fmt.Sprintf("%s/robots.txt", topLevelUrl)
	NodeDB.Link = topLevelUrl
	cNode := []Node{ Node{}}
	NodeDB.ChildNodes = cNode

	if TopLevelUrl != "" {
		sitemaps := ProcessRobotsFile(robotsFile)

		for _, sitemap := range sitemaps {
			ProcessSiteMapFile(sitemap)
		}

		CreateReport(reportFile1)

		BuildHierachyTree()
		CreateTreeReport(reportFile2)
	}
}

func ProcessRobotsFile(robotsFile string) []string {
	var sitemaps []string

	if len(robotsFile) != 0 {
		robotsPage := GetWebText(robotsFile)
		sitemaps = ParseRobotsFile(robotsPage)
	}

	return sitemaps
}

func ParseRobotsFile(robotsFile string) []string {
	var sitemaps []string

	sitemapRegexLinePattern, err := regexp.Compile(`Sitemap: * http.*\.xml`) // (`Sitemap: * http.*\.xml`)
	sitemapRegexPattern, err := regexp.Compile(`http.*\.xml`)

	matches := sitemapRegexLinePattern.FindAllString(robotsFile, -1)

	if err == nil {
		for _, match := range matches {
			sms := sitemapRegexPattern.FindAllString(match, -1)
			if len(sms) == 1 {
				sm := strings.TrimSpace(sms[0])
				sitemaps = append(sitemaps, sm)
			}
		}
	} else {
		log.Fatal(err)
	}

	return sitemaps
}

func ProcessSiteMapFile(siteMapFile string) {
	if len(siteMapFile) != 0 {
		siteMapXmlPage := GetWebText(siteMapFile)
		siteMapXml := ParseSiteMapFile(siteMapXmlPage)

		if sitemaps, ok := siteMapXml["sitemap"]; ok {
			for _, sitemap := range sitemaps {
				ProcessSiteMapFile(sitemap)
			}
		}
	}
}

func ParseSiteMapFile(siteMapPage string) map[string][]string {
	links := make(map[string][]string)

	sitemapValues := []string{""}
	weblinkValues := []string{""}
	links["sitemap"] = sitemapValues
	links["link"] = weblinkValues

	parsedLinks := XmlParse([]byte(siteMapPage))

	for _, plink := range parsedLinks {
		var link = strings.TrimSpace(plink)

		if strings.HasSuffix(link, ".xml") {
			// duplicate check?
			sitemapValues = append(sitemapValues, link)
		} else {
			// duplicate check?
			// weblinkValues = append(weblinkValues, link)
			LinkDB = append(LinkDB, link)
		}
	}

	// fmt.Println("link2s = ", link2s)
	return links
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
		result = string(contents)
	}
	return result
}

func CreateReport(reportFile string) {
	file, err := os.Create(reportFile)
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()

	for _, weblink := range LinkDB {
		fmt.Fprintln(file, weblink)
	}
}

func BuildHierachyTree() {
	var activeNode Node
	activeNode = NodeDB

	for _, weblink := range LinkDB {
		if strings.Contains(weblink, TopLevelUrl) {
			var weblinkWithoutTopLevelUrl = strings.Replace(weblink, TopLevelUrl, "", -1)
			var weblinkTokens = strings.Split(weblinkWithoutTopLevelUrl, "/")

			if len(weblinkTokens) > 0 {
				activeNode = NodeDB
				for _, token := range weblinkTokens {
					trimmedToken := strings.TrimSpace(token)
					if trimmedToken != "" {
						childNodeExists := false
						var childNode Node

						if len(activeNode.ChildNodes) > 0 {
							for _, cnode := range activeNode.ChildNodes {
								if cnode.Link == trimmedToken {
									childNode = cnode
									childNodeExists = true
									break
								}
							}
						}

						if (!childNodeExists) {
							// Add new node
							childNode = Node{}
							childNode.Link = trimmedToken
							activeNode.ChildNodes = append(activeNode.ChildNodes, childNode)
						}
						activeNode = childNode
					}
				}
			}

		}

	}
	fmt.Printf("DBG1: LinkDB size = %d\n", len(LinkDB))
	fmt.Printf("DBG1: NodeDB size = %d\n", len(NodeDB.ChildNodes))

}

func CreateTreeReport(reportFile2 string) {
	var activeNode Node
	activeNode = NodeDB

	var displayStr []string
	var level int
	level = 1
	displayStr = append(displayStr, activeNode.Link)
	level = level + 1
	displayStr = PrintNode(activeNode, displayStr, level)

	fmt.Println("displayStr = ", len(displayStr))

	file, err := os.Create(reportFile2)
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()

	for _, str := range displayStr {
		fmt.Fprintln(file, str)
	}
}

func PrintNode(activeNode Node, displayStr []string, level int) []string {
	var tabstr string
	tabstr = ""
	for ii := 0; ii < level; ii++ {
		tabstr = tabstr + "    "
	}

	level = level + 1

	for _, childNode := range activeNode.ChildNodes {
		displayStr = append(displayStr, tabstr+childNode.Link)
		displayStr = PrintNode(childNode, displayStr, level)
	}
	fmt.Println("displayStr = ", len(displayStr))

	return displayStr
}

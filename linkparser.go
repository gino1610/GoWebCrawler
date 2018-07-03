package main

import (
	"strings"
  )

var webLinkRegex string = "href=\"[a-zA-Z./:&\\d_-]+\""
var goodUrls []string
var badUrls []string
var otherUrls []string
var externalUrls []string
var exceptions []string

func ParseLinks(page Page, url string) {
	
}

func IsExternalUrl(url string) bool {
	return false
}

func IsAWebPage(foundHref string) bool {
	firstIndex := strings.IndexAny(foundHref, "javascript:")
	if (firstIndex == 0) {
		return false
	}

	length := len(foundHref)
	lastIndex := strings.LastIndexAny(foundHref, ".")
	if (length > 0 && lastIndex != -1) {
		splits :=  strings.Split(foundHref, ".")
		extension := splits[len(splits)-1]

		switch extension {
		case "jpg":
			return false
		case "css":
			return false
		default:
			return true
		}

	} else {
		return true;
	}
}
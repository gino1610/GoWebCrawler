package main

type urlRepo struct {
	urlList 	[]string	
}

func (list urlRepo) Add (url string) {
	list.urlList = append(list.urlList, url)
}
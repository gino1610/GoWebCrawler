# GoWebCrawler
Web Crawler in Go lang

Note: As I am new to both Web crawlers and Go lang, I have relied on various web resources to learn these 2 items. More specifically, this go web crawler example is heavily inspired from https://github.com/NeelBhatt/WebCrawler for the core functionality


Instructions to build and run
- cd to ..../GoWebCrawler
- Run the below command
	go run *.go

	
Notes Update
- The current code is work-in-progress wherein, I could send a HttpRequest and get the response back. The following items are pending at this time
Parse the Webpage response and extract links (this will go in ParseLinks function in the file: LinkParser.go)
- Following the links in the first page and then crawling those links iteratively. (this will go in crawlPage function in the crawler.go file)

- Once the crawling is complete (based on a user stopping the process or limiting to number of url page links or some metric ---- I have not yet thought through on this), a report would be generated for now (a sample output file is attached to this email). Later on, the results could be either stored in a flat file or a database so that the user could query and get specific results).

- And I will continue to fix/update/refine the application.
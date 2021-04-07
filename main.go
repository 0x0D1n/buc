package main

import (
	"fmt"
	"flag"
	_"os"
	_"strings"
)


func main() {

	var url string
	flag.StringVar(&url, "u", "", "URL for dead links checker")

	var filename string
	flag.StringVar(&filename, "f", "", "file containing URLs to test")

	var writeToFile string 
	//Let the user choose the HTTP code to store in the file ?
	flag.StringVar(&writeToFile, "w", "", "filename for storage of != 200 HTTP Error codes URLs")

	flag.Parse()

	var params string


	if url != "" {
		params = "u"
	}else if filename != "" {
		params = "f"
	}else{
		flag.PrintDefaults()
	}

	var interestURLs []string

	if params == "f" {
		var tmpInterestURLs []string
		urls, _ := readUrlsFromFile(filename)
		for i := 0 ; i < len(urls) ; i++ {
			tmpInterestURLs = checkDeadLinks(urls[i])
			for x := range tmpInterestURLs{
				interestURLs = append(interestURLs, tmpInterestURLs[x])
			}
		}
	}

	if params == "u" {
		interestURLs = checkDeadLinks(url)
	}

	if writeToFile != "" {
		fmt.Println("[+] Written results to file > ", writeToFile)
		writeUrlsToFile(writeToFile, interestURLs)
	}

}

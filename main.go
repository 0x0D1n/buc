package main

import (
	"fmt"
	"flag"
	"os"
	_"strings"
)


func main() {

	var url string
	flag.StringVar(&url, "u", "", "URL for dead links checker")

	var filename string
	flag.StringVar(&filename, "f", "", "file containing URLs to test")

	var writeToFile string 
	flag.StringVar(&writeToFile, "w", "", "Filename for storage")

	flag.Parse()

	var params string

	if url != ""{
		params = "u"
		fmt.Println("URL")
	}else if filename != ""{
		params = "f"
		fmt.Println("File")
	}else{
		params = ""
		flag.PrintDefaults()
		os.Exit(0)
	}

	if params == "f" {
		urls, _ := readUrlsFromFile(filename)
		for i := 0 ; i < len(urls) ; i++ {
			checkDeadLinks(urls[i])
		}
	}

	if params == "u" {
		checkDeadLinks(url)
	}

}

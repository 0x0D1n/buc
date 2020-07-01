package main

import (
	"fmt"
	"flag"
	"os"
	"strings"
)


func main() {

	var url string
	flag.StringVar(&url, "u", "", "Url for dead links checker")

	var writeToFile string 
	flag.StringVar(&writeToFile, "w", "", "Filename for storage")

	flag.Parse()

	if url == "" {
		flag.PrintDefaults()
		os.Exit(0)
	}

	if strings.HasPrefix(url, "http") == false {
		fmt.Println("Please enter a scheme like http before the url")
		os.Exit(0)
	}
	
	checkDeathLinks(url)

}

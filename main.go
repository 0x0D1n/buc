package main

import (
	_ "fmt"
	"flag"
	"os"
)



func main() {

	var url string
	flag.StringVar(&url, "u", "", "Url for dead links checker")

	var writeToFile string 
	flag.StringVar(&writeToFile, "w", "", "Filename for storage")

	flag.Parse()

	if url == ""{
		flag.PrintDefaults()
		os.Exit(0)
	}
	
	checkDeathLinks(url)

}

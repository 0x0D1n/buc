package main

import (
	"fmt"
	_ "flag"
	"os"
)



func main() {
	
	//var url string
	//urlPtr := flag.StringVar(&url, "u", "url", "https://example.com")

	//flag.Parse()
	
	fmt.Println("Arg 0 : ", os.Args[1])

	if url := os.Args[1]; url != "" {
		retrieveUrls(url)
	}

}

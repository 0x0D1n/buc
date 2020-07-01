package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"github.com/mvdan/xurls"
	"strings"
	"log"
)

func removeDupUrls(urls []string) []string {
	
	encountered := map[string]bool{}
	finalList := []string{}
	
	for v := range urls {
		if encountered[urls[v]] == true {
			
		} else {
			encountered[urls[v]] = true
			finalList = append(finalList, urls[v])
		}
	}

	return finalList
}


func retrieveUrls(sourceCode string) []string{ 
	
	rxStrict := xurls.Strict()
	rxRelaxed := xurls.Relaxed()
	urls := []string{}
	urls = rxStrict.FindAllString(sourceCode, -1)
	urls = append(urls, rxRelaxed.FindString(sourceCode))
	urls = removeDupUrls(urls)
	return urls
}


func checkDeathLinks(url string) {
	
	resp, err := http.Get(url)
	
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)
	links := retrieveUrls(bodyString)
	
	//TODO - HANDLE THE NEWS LINKS WE GOT
	fmt.Print(strings.Join(links[:], "\n"))	

}

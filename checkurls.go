package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"github.com/mvdan/xurls"
	_ "strings"
	"log"
	_ "bufio"
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

func isLinkDead(urls []string) {
	
	for _, v := range urls {
		resp, err := http.Get(v)
		if err != nil {
			fmt.Println(err)
		}

		defer resp.Body.Close()
		
		beautifyOutput(v, resp.StatusCode, http.StatusText(resp.StatusCode))
	}
}


//Final function doing everything
func checkDeadLinks(url string) {
	
	resp, err := http.Get(url)
	
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)
	links := retrieveUrls(bodyString)
	
	//fmt.Print(strings.Join(links[:], "\n"))	
	isLinkDead(links)

}

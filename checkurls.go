package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"github.com/mvdan/xurls"
	"strings"
	"log"
	_ "bufio"
)

type urlStatus struct {
	url		string
	status	bool
}

func removeDupUrls(urls []string) []string {
	
	encountered := map[string]bool{}
	finalList := []string{}
	
	for v := range urls {
		if encountered[urls[v]] == true {
		
		}else {
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
	c := make(chan urlStatus)
	for _, v := range urls {
		go makeHttpRequest(v, c)
	}
	result := make([]urlStatus, len(urls))
	for i, _ := range result {
		result[i] = <-c
		if result[i].status {
			fmt.Println(result[i].url, "is up.")
		} else {
			fmt.Println(result[i].url, "is down !!")
		}
	}
		
	//beautifyOutput(v, resp.StatusCode, http.StatusText(resp.StatusCode))
}


func makeHttpRequest(url string, c chan urlStatus){
	_, err := http.Get(url)
	if err != nil {
		//HTTP Failure
		c <- urlStatus{url, false}
	}else {
		//HTTP Success
		c <- urlStatus{url, true}
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
	
	fmt.Print(strings.Join(links[:], "\n"))	
	//fmt.Println()
	//fmt.Println("------------------------")
	//fmt.Println()
	isLinkDead(links)

}

package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	_"github.com/mvdan/xurls"
	"mvdan.cc/xurls/v2"
	_ "strings"
	"log"
	_ "bufio"
	"strconv"
)

type urlStatus struct {
	url	string
	status	bool
	statusCode int
	statusText string
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
	/*
	xurls gathers URLs containing ':' -> Leading to wrong Urls being gathered
	Needs to be done manually - Regex
	*/
	rxStrict := xurls.Strict()
	rxRelaxed := xurls.Relaxed()
	urls := []string{}
	urls = rxStrict.FindAllString(sourceCode, -1)
	urls = append(urls, rxRelaxed.FindString(sourceCode))
	urls = removeDupUrls(urls)
	return urls
}

func isLinkDead(urls []string) []string {
	c := make(chan urlStatus)
	var interestURLs []string
	var fmtStringURL string

	for _, v := range urls {
		go makeHttpRequest(v, c)
	}
	result := make([]urlStatus, len(urls))
	for i, _ := range result {
		result[i] = <-c
		if result[i].statusCode != 0 {
			if result[i].status {
				//fmt.Println(result[i].url, "is up and status code is ", result[i].statusCode)
				beautifyOutput(result[i].url, result[i].statusCode, result[i].statusText)
				//For storage to file of the interessting HTTP status codes
				//Better way then != 200 ? Should be enough for the moment
				if result[i].statusCode != 200 {
					fmtStringURL = string(result[i].url) + "," + strconv.Itoa(result[i].statusCode)
					interestURLs = append(interestURLs, fmtStringURL)
					fmt.Println(interestURLs)
				}
			}
		}
	}
	return interestURLs
}


func makeHttpRequest(url string, c chan urlStatus){
	resp, err := http.Get(url)
	var statusCode int
	var statusText string
	if err != nil {
		//HTTP Failure
		c <- urlStatus{url, false, statusCode, statusText}
		//fmt.Println("Failure : ",url)
	}else {
		//HTTP Success
		statusCode := resp.StatusCode
		statusText := http.StatusText(resp.StatusCode)
		c <- urlStatus{url, true, statusCode, statusText}
	}

}


//Final function doing everything
func checkDeadLinks(url string) []string{
	
	resp, err := http.Get(url)
	
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)
	links := retrieveUrls(bodyString)
	
	//fmt.Print(strings.Join(links[:], "\n"))	
	interestURLs := isLinkDead(links)

	return interestURLs
}

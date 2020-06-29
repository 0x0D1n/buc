package main


import (
	"fmt"
	"net/http"
	"io/ioutil"
)


func retrieveUrls(url string) {
	
	resp, err := http.Get(url)
	
	if err != nil {
		fmt.Errorf("Can't /GET", url, err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(body)

}

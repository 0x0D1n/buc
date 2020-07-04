package main

import (

	"fmt"
	"github.com/logrusorgru/aurora"
)


const (
	InformativeCode = 100	// Yellow ?
	SuccessCode 	= 200	// Green ?
	RedirectionCode = 300	// Magenta ?
	ClientErrorCode = 400	// Red ?
	ServerErrorCode = 500	// Red ?
)


func beautifyOutput(url string, statuscode int, statustext string) {
	
	if statuscode >= RedirectionCode {
		fmt.Println(url, " -> ", aurora.Red(statuscode), aurora.Red(statustext))
	}else {
		fmt.Println(url, " -> ", aurora.Green(statuscode), aurora.Green(statustext))
	}
}

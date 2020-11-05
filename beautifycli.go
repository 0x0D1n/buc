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
	
	switch {
		case statuscode >= 100 && statuscode <= 199:
			fmt.Println("[+] ", url, "  ", aurora.Yellow(statuscode), aurora.Yellow(statustext))
		case statuscode >= 200 && statuscode <= 299:
			fmt.Println("[+] ", url, "  ", aurora.Green(statuscode), aurora.Green(statustext))
		case statuscode >= 300 && statuscode <= 399:
			fmt.Println("[+] ", url, "  ", aurora.Magenta(statuscode), aurora.Magenta(statustext))
		case statuscode >= 400 && statuscode <= 499:
			fmt.Println("[+] ", url, "  ", aurora.Red(statuscode), aurora.Red(statustext))
		case statuscode >= 500:
			fmt.Println("[+] ", url, "  ", aurora.Red(statuscode), aurora.Red(statustext))
	}
}





package main


import (
	"bufio"
	"fmt"
	"os"
)

//https://stackoverflow.com/questions/5884154/read-text-file-into-string-array-and-write
func readUrlsFromFile(path string) ([]string, error){
	f, err := os.Open(path)

	if err != nil {
		//return nil, err
		fmt.Println("[+] Error occured during file handling ...")
		os.Exit(0)
	}

	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}


func writeUrlsToFile(filename string, lines []string) error {

	f, err := os.Create(filename)
	
	if err != nil {
		return err
	}

	defer f.Close()

	w := bufio.NewWriter(f)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}

	return w.Flush()
}


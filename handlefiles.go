package main


import (
	"bufio"
	"fmt"
	"os"
)


func readUrlsFromFile(path string) ([]string, error) {
	f, err := os.Open(path)

	if err != nil {
		return nil, err
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


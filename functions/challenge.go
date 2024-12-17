package main

import (
	"fmt"
	"net/http"
)

func main() {

	cType, err := contentType("https://www.google.com")
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
	} else {
		fmt.Println(cType)
	}
}
func contentType(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("Error fetching URL %s: %s", url, err)
	}
	defer resp.Body.Close()
	return resp.Header.Get("Content-Type"), nil

}

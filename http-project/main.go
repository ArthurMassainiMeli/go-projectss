package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Printf("Error fetching url\nError: %s\n", err.Error())
		os.Exit(1)
	}

	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	_, err = io.Copy(os.Stdout, resp.Body)
	if err != nil {
		fmt.Printf("Error getting body url\nError: %s\n", err.Error())
		os.Exit(1)
	}
}

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	// Get request
	res, err := http.Get("http://localhost:1323/prefectures-json")
	if err != nil {
		log.Fatal(err)
	}

	// header
	fmt.Printf("[status] %d\n", res.StatusCode)
	for k, v := range res.Header {
		fmt.Print("[header] " + k)
		fmt.Println(": " + strings.Join(v, ","))
	}

	// body
	defer res.Body.Close()
	body, error := ioutil.ReadAll(res.Body)
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println("[body] " + string(body))
}

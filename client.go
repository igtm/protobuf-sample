package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/igtm/protobuf-sample/pb"
)

func main() {
	// Get request
	res, err := http.Get("http://localhost:1323/prefectures")
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
	// fmt.Println("[body] " + string(body))

	pref := &pb.GetPrefecturesResponse{}
	if err := proto.Unmarshal(body, pref); err != nil {
		log.Fatalln("Failed to parse prefectures:", err)
	}

	for _, p := range pref.Prefectures {
		fmt.Printf("%v, %v, %v\n", p.Id, p.Name, p.Romaji)
	}
}

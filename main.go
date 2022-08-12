package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// No idea why this isnt working. i'm new to go pls halp :c
	output := flag.String("i", "", "IP address of you want to get information about")
	flag.Parse()
	println(*output)
	if *output == "" {
		flag.Usage()
		return
	}
	fmt.Println("checking IP address", *output, "...")

	response, err := http.Get("http://ip-api.com/json/" + *output)
	if err != nil {
		log.Fatal(err)
	}
	// structure of response
	type Response struct {
		Status      string  `json:"status"`
		Message     string  `json:"message"`
		Country     string  `json:"country"`
		CountryCode string  `json:"countryCode"`
		Region      string  `json:"region"`
		RegionName  string  `json:"regionName"`
		City        string  `json:"city"`
		Zip         string  `json:"zip"`
		Lat         float64 `json:"lat"`
		Lon         float64 `json:"lon"`
		Timezone    string  `json:"timezone"`
		Offset      float64 `json:"offset"`
		Isp         string  `json:"isp"`
		Org         string  `json:"org"`
	}

	var responseData Response
	err = json.NewDecoder(response.Body).Decode(&responseData)
	responseDataJSON, err := json.MarshalIndent(responseData, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseDataJSON))
}

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Country struct {
	Name       string `json:"name"`
	Population int    `json:"population"`
	Capital    string `json:"capital"`
}

func main() {
	// Step 1: Define the endpoint
	endpoint := "https://api.sampleapis.com/countries/countries"

	// Step 2: Define struct to store country information
	var countries []Country

	// Step 3: Make HTTP GET request and unmarshal response
	resp, err := http.Get(endpoint)
	if err != nil {
		log.Fatal("Error sending request:", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response body:", err)
	}

	err = json.Unmarshal(body, &countries)
	if err != nil {
		log.Fatal("Error unmarshaling response body:", err)
	}

	// Step 4: Print country information
	for _, country := range countries {
		fmt.Println("Name:", country.Name)
		fmt.Println("Population:", country.Population)
		fmt.Println("Capital:", country.Capital)
		fmt.Println("")
	}
}

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
)

type Media struct {
	Flag         string `json:"flag"`
	Emblem       string `json:"emblem"`
	Orthographic string `json:"orthographic"`
}

type CountryInfo struct {
	Abbreviation string `json:"abbreviation"`
	Capital      string `json:"capital"`
	Currency     string `json:"currency"`
	Name         string `json:"name"`
	Phone        string `json:"phone"`
	Population   int    `json:"population"`
	Media        Media  `json:"media"`
	ID           int    `json:"id"`
}

// Sort countries by population
func sortByPopulation(countries []CountryInfo) {
	sort.Slice(countries, func(i, j int) bool {
		return countries[i].Population > countries[j].Population
	})
}

// Sort countries by name
func sortByName(countries []CountryInfo) {
	sort.Slice(countries, func(i, j int) bool {
		return countries[i].Name < countries[j].Name
	})
}

// Define and implement a function to get a country by its abbreviation
func getCountryByAbbreviation(countries []CountryInfo, abbreviation string) CountryInfo {
	// TODO: Implement the logic to find and return the country with the given abbreviation
	// Iterate over the countries slice and check the abbreviation field for a match
	// Return the country when found, or a default CountryInfo if not found

	// Placeholder code:
	var country CountryInfo
	return country
}

func main() {
	// Step 1: Use the http.Get function to get the response from the API endpoint
	response, err :=
		http.Get("https://api.sampleapis.com/countries/countries")

	if err != nil {
		fmt.Println(err)
		return
	}

	// Step 2: Read all the data from the response
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Step 3: Create an empty slice of CountryInfo
	var countries []CountryInfo

	// Unmarshal the data and print an error if it exists
	err = json.Unmarshal(data, &countries)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Sort countries by population and print them

	// Sort countries by name and print them

	// Define and implement a function to get a country by its abbreviation

	// Use the above function to get and print the details of the country with abbreviation 'US'

	// Print the media URLs for each country
}

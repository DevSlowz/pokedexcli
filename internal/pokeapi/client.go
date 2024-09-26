package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type LocationAreaList struct {
	Next     string  `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

// type Config struct {
// 	Next     *string
// 	Previous *string
// }

var Base_URL = "https://pokeapi.co/api/v2/location-area/"

func GetInfo(cnfg *Config) ([]string, error) {
	// resp, err := http.Get("https://pokeapi.co/api/v2/location-area/")
	resp, err := http.Get(*cnfg.Previous)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// Parse JSON
	var locationList LocationAreaList
	err = json.Unmarshal(body, &locationList)
	if err != nil {
		log.Fatalln(err)
	}

	var linkConfg Config
	linkConfg.Next = &locationList.Next
	//test := locationList.Next
	// fmt.Println("Next:", linkConfg.Next)

	linkConfg.Previous = &url
	// fmt.Println("Previous:", linkConfg.Previous)

	// fmt.Println("Results:")
	// for _, result := range locationList.Results {
	// 	fmt.Printf("  - %s (%s)\n", result.Name, result.URL)
	// }
	var results []string
	for _, result := range locationList.Results {
		results = append(results, result.Name)
	}

	// My Question now is - How do I get it to go to the next page?
	// or
	// How do store the Next and Previous links in my struct?
	return results, nil
}

func mapf(cnfg *Config) error {

	var names []string
	if cnfg.Previous == nil || *cnfg.Previous == "" {
		results, _, _ := GetInfo(Base_URL)
		names = results
	}

	for _, result := range names {
		fmt.Printf("(%s)\n", result)
	}
	return nil
}

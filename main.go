package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type characters struct {
	Character []character `json:"results"`
}

type character struct {
	Name string `json:"name"`
	Status string `json:"status"`
	Species string `json:"species"`
	Gender string `json:"gender"`
}

func getCharacters(url string)(characters, error) {
	c := characters{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return c, err
	}
	req.Header.Set("User-Agent", "rickandmorty-tutorial")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return c, err
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return c, err
	}

	err = json.Unmarshal(body, &c)
	if err != nil {
		return c, err
	}

	return c, nil
}

func main() {
	url := "https://rickandmortyapi.com/api/character"

	characters, err := getCharacters(url)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range characters.Character {
		fmt.Printf("Name - %s | Status - %s | Species - %s | Gender - %s\n" , v.Name, v.Status, v.Species, v.Gender)
		}
}
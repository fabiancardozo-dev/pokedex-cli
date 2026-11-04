package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type Location struct {
	Count    int       `json:"count"`
	Next     string    `json:"next"`
	Previous *string   `json:"previous"`
	Results  []Results `json:"results"`
}

type Results struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func commandMap(cfg *config) error {
	location, err := getLocations(cfg)
	if err != nil {
		return err
	}

	for _, location := range location.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func getLocations(cfg *config) (Location, error) {
	res, err := http.Get(cfg.currentURL)
	if err != nil {
		return Location{}, err
	}

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode > 299 {
		return Location{}, errors.New("Request failed with status code: " + string(res.StatusCode))
	}
	if err != nil {
		return Location{}, err
	}

	var loc Location
	err = json.Unmarshal(body, &loc)
	if err != nil {
		return Location{}, err
	}

	cfg.previousURL = cfg.currentURL
	cfg.currentURL = loc.Next

	return loc, nil
}

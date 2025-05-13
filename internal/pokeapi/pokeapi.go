package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/marekmchl/pokedexcli/internal/pokecache"
)

type Map struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}

func GetMap(url string, cache pokecache.Cache) (Map, error) {
	jsonData := []byte{}
	cashed, found := cache.Get(url)
	if found {
		// fmt.Println("(cached)")
		jsonData = cashed
	} else {
		// fmt.Println("(downloaded)")
		res, err := http.Get(url)
		if err != nil {
			return Map{}, fmt.Errorf("failed: pokapi.go - GetMap: %v", err)
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if res.StatusCode > 299 {
			return Map{}, fmt.Errorf("failed: pokapi.go - GetMap: with status code: %d and\nbody: %s\n", res.StatusCode, body)
		}
		if err != nil {
			return Map{}, fmt.Errorf("failed: pokapi.go - GetMap: %v", err)
		}
		cache.Add(url, body)
		jsonData = body
	}

	data := Map{}
	if err := json.Unmarshal(jsonData, &data); err != nil {
		return Map{}, fmt.Errorf("failed: pokapi.go - GetMap: %v", err)
	}

	return data, nil
}

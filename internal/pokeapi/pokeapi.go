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

type Location struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func GetPokemons(location string) ([]string, error) {
	url := "https://pokeapi.co/api/v2/location-area/" + location + "/"
	res, err := http.Get(url)
	if err != nil {
		return []string{}, fmt.Errorf("failed: pokapi.go - GetPokemons - get: %v", err)
	}
	defer res.Body.Close()

	rawJson, err := io.ReadAll(res.Body)
	if err != nil {
		return []string{}, fmt.Errorf("failed: pokapi.go - GetPokemons - read: %v", err)
	}

	locationMap := &Location{}
	if err := json.Unmarshal(rawJson, locationMap); err != nil {
		return []string{}, fmt.Errorf("failed: pokapi.go - GetPokemons - unmarshal: %v", err)
	}

	pokemons := []string{}
	for _, pokemonEncounter := range locationMap.PokemonEncounters {
		pokemons = append(pokemons, pokemonEncounter.Pokemon.Name)
	}

	return pokemons, nil
}

package pokeapi

import (
	"fmt"
	"math"
	"math/rand/v2"
)

func (c *Client) GetPokemon(name string) (PokemonResponse, error) {
	var response PokemonResponse
	url := "pokemon/" + name
	err := c.fetchData(&url, &response)
	if err != nil {
		return PokemonResponse{}, err
	}

	return response, nil
}

func CatchPokemon(pokeData PokemonResponse) bool {
	fmt.Printf("Throwing a Pokeball at %s...\n", pokeData.Name)
	catchResistance := math.Log(1.0 + float64(pokeData.BaseExperience)) // between 0 and ~6.4
	randomNum := 1.6 + rand.Float64()*5.8
	//log.Printf("Catch Resistance: %.2f, Random number: %.2f", catchResistance, randomNum)
	if randomNum > catchResistance {
		fmt.Printf("%s was caught!\n", pokeData.Name)
	} else {
		fmt.Printf("%s escaped!\n", pokeData.Name)
	}
	return randomNum > catchResistance
}

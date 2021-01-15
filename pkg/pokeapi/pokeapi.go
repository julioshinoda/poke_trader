package pokeapi

import (
	apiv2 "github.com/mtslzr/pokeapi-go"
	apiv2Struct "github.com/mtslzr/pokeapi-go/structs"
)

func GetPokemon(pokemon string) (apiv2Struct.Pokemon, error) {
	return apiv2.Pokemon(pokemon)
}

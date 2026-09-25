package main

import "github.com/DorothyDanger/Pokedex/internal/pokecache"

/*
Contains struct of Next and Previous URLs
This is to paginate through location areas
*/

type config struct {
	next     *string
	previous *string
	cache    *pokecache.Cache
	pokedex  map[string]pokemon
}

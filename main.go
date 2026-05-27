package main

import "github.com/Sed-Miyuki/PokedexCli/internal/pokeapi"

type config struct{
	pokeapiClient pokeapi.Client
	nextLocationAreaURL *string
	prevLocationAreaURL *string
}

func main(){
	cfg:=&config{
		pokeapiClient:pokeapi.NewClient(),
	}
	startRepl(cfg)
}

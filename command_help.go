package main

import "fmt"

func callbackHelp(cfg *config) error{
	println("Hello Welcome to Pokedex help menu!!!")
	println("Here are your available command!!!")
	availableCommands:=getCommands()
	for _,cmd:=range(availableCommands){
		fmt.Printf(" - %s:%s\n",cmd.name,cmd.description)
	}
	return nil
}
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct{
	name 		string
	description string
	callback 	func(*config) error
}

func startRepl(cfg *config){
	reader:=bufio.NewScanner(os.Stdin)
	for{
		fmt.Print("Pokedex >")
		text:=reader.Text()
		words:=cleanInput(text)
		if len(words)==0{
			continue
		}
		commandName:=words[0]
		availableCommands:=getCommands()
		command,ok:=availableCommands[commandName]
		if !ok{
			fmt.Println("invalid command")
			continue
		}
		err:=command.callback(cfg)
		if err!=nil{
			fmt.Println(err)
		}
	}
}

func cleanInput(str string) []string{
	lowered:=strings.ToLower(str)
	words:=strings.Fields(lowered)
	return  words
}

func getCommands() map[string]cliCommand{
	return map[string]cliCommand{
		"help":{
			name: 			"help",
			description: 	"Prints the help menu",
			callback:  		callbackHelp,
		},
		"exit":{
			name: 			"exit",
			description: 	"Turn off the pokedex",
			callback:  		callbackExit,
		},
		"map":{
			name: 			"map",
			description: 	"List next page Loaction areas",
			callback:  		callbackMap ,
		},
		"mapb":{
			name: 			"mapb",
			description: 	"List previous page Loaction areas",
			callback:  		callbackMapb ,
		},
	}
}
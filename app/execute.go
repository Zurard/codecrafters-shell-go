package main 

import (
	"fmt"
	"strings"	
	"errors"
)


func execute(cmd string , args []string) { 
	switch cmd {
	// exit command stops the shell from taking more prompts  
	case "exit" : exit()

	case "echo" : echo(args)

	case "type" : IsValidCommmand(args) 
}

} 

func exit() { 	
	toExit = true  
}

func echo(args []string){
	output := strings.Join(args," ")
	fmt.Println(output)
}

func IsValidCommmand(args []string) {
	cmd := strings.Join(args," ")
	_ , found := commands[cmd] 
	if found { 
		fmt.Println(cmd + " is a shell built in")
		} else {
		fmt.Println( errors.New(cmd + ": not found"))
	}
}




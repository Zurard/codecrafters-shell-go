package main 

import (
	"fmt"
	"strings"	
)


func execute(cmd string , args []string) { 
	switch cmd {
	// exit command stops the shell from taking more prompts  
	case "exit" : exit()

	case "echo" : echo(args)
}
} 

func exit() { 	
	toExit = true  
}

func echo(args []string){
	output := strings.Join(args," ")
	fmt.Println(output)
}






// // need to parse the cmd 

package main 

import (
	"strings"
)

var toExit = false

func parse(input string){
	text_arr := strings.Fields(input)

	cmd := text_arr[0]
	args := text_arr[1:]
	
	if cmd == "exit" { 
		exit()
	} else if cmd == "echo" {
		echo(args)
	} else if cmd == "type" {
		execType(args)
	} else if cmd == "jobs"{
		if len(args) > 1 {
			// then we need to do something but for now nothing 
			// runJobs(cmd , args[1:len(args)-1])
		}
	} else if len(cmd) > 0 && len(args)-1 >= 0 && args[len(args)-1] == "&" {
		if len(args)- 1 > 0 {
			runJobs(cmd , args[1:len(args)-1])
			} 
	} else { 
		runProgram(cmd, args)}
	}



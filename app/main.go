package main

// fmt is a popular package in go for formatting text
import (
	"bufio" 
	"strings"
	"errors"
	"os"
	"fmt")

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print

// this are the lest of commands we can execute in our code 
var commands = make(map[string] struct{})

func parse(cmd string) (string,error) {
	_, found := commands[cmd]
	// fmt.Println(found)
	if !found {
		err := cmd + ": " + "command not found"  
		return "",errors.New(err + "\n" )
	}
	return cmd , nil 
}

func read() {
	// reader object
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("$ ")
	text,_ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	for text == "" {
		fmt.Print("$ ")
		text,_ = reader.ReadString('\n')
		text = strings.TrimSpace(text)
	}  
	_, err := parse(text)
	// so basically if error is not null 
	if err != nil {
		// here we will get the error created in parse function as in stage 3 : err is {cmd} command not found 
		fmt.Print(err)
	}
}

func main() {
	// keep taking commands until and unless we exit the loop 
	for true {
		read()
	}	
}

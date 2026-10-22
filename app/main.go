package main

// fmt is a popular package in go for formatting text
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// if a tempraray path is also passed then we need to check if output exist in that path or not in all the directories of that path
var currPath = os.Getenv("PATH")
var IsCurrPath = false

// will keep taking input for us 
func read() {
	// reader object
	reader := bufio.NewReader(os.Stdin)
	// fmt.Println(currPath)
	if len(currPath) > 0 {
		IsCurrPath = true
	}
	fmt.Print("$ ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)

	for text == "" {
		fmt.Print("$ ")
		text, _ = reader.ReadString('\n')
		text = strings.TrimSpace(text)
	}
	// PARSE THE TEXT ENTERED
	parse(text)
}

func main() {
	// keep taking commands until and unless we exit the loop
	for true {
		read()
		if toExit == true {
			return
		}
	}
}

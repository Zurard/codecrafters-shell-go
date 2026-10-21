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
	// need to extract the command and its arguments
	text_arr := strings.Fields(text)

	cmd := text_arr[0]
	args := text_arr[1:]

	// _, ifExist := commands[cmd]

	// var err error = nil

	// if !ifExist {
	// 	err = fmt.Errorf("%s: command not found\n", cmd)
	// }
	// // so basically if error is not null
	// if err != nil {
	// 	// here we will get the error created in parse function as in stage 3 : err is {cmd} command not found
	// 	fmt.Print(err)
	// }
	execute(cmd, args)

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

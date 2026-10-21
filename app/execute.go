package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func execute(cmd string, args []string) {
	switch cmd {
	// exit command stops the shell from taking more prompts
	case "exit":
		exit()

	case "echo":
		echo(args)

	case "type":
		execType(args)

	default:
		runProgram(cmd, args)
	}

}

func exit() {
	toExit = true
}

func echo(args []string) {
	output := strings.Join(args, " ")
	fmt.Println(output)
}

func execType(args []string) {
	output := strings.Join(args, " ")
	_, found := commands[output]

	// the path would look something like this :
	// /usr/bin:/usr/local/bin:/home/shree/.local/bin:/home/shree/.bun/bin:/home/shree/.config/Code/User/globalStorage/github.copilot-chat/debugCommand:/home/shree/.config/Code/User/globalStorage/github.copilot-chat/copilotCli:/home/shree/.local/bin:/home/shree/.local/bin:/home/shree/.bun/bin:/home/shree/.nvm/versions/node/v24.11.1/bin:/home/shree/.local/share/pnpm:/home/shree/.cargo/bin:/home/shree/.local/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/games:/usr/local/games:/snap/bin:/snap/bin:/home/shree/ida-free-pc-9.2:/home/shree/ida-free-pc-9.2

	if found {
		fmt.Println(output + " is a shell builtin")
	} else if IsCurrPath {
		path, err := exec.LookPath(output)
		if err != nil {
			fmt.Println(errors.New(output + ": not found"))
		} else {
			fmt.Println(output + " is " + path)
		}

	}
}

func runProgram(cmd string, args []string) {
	// fmt.Println("curr path is :", currPath )
	// fmt.Println("program which needs to be executed" ,cmd )
	// fmt.Println("arguments are : ", strings.Join(args, " "))

	// path , err := exec.LookPath(cmd)
	// fmt.Println("path is : " , path , "error is :" , err )
	_, err := exec.LookPath(cmd)
	if err == nil {
		output := exec.Command(cmd, args[0:]...)
		output.Stdout = os.Stdout
		output.Stderr = os.Stderr
		output.Run()
	} else {
		// fmt.Println(len(cmd))
		fmt.Println(cmd,": command not found")
	}

}

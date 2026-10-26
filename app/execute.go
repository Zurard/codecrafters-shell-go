package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

var processNo = 0  

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
	_, err := exec.LookPath(cmd)
	if err == nil {
		output := exec.Command(cmd, args[0:]...)
		output.Stdout = os.Stdout
		output.Stderr = os.Stderr
		output.Run()
	} else {
		// fmt.Println(len(cmd))
		fmt.Println(cmd + ": command not found")
	}

}


// helper function for debugging 

func getLinuxProcessPath(pid int) (string , error ) {
	procPath := fmt.Sprintf("/proc/%d/exe", pid) 
	path , err := os.Readlink(procPath)

	if err != nil {
		return "" , err 
	}
	return path , nil 
}

func runJobs(command string , args []string) { 
	// now i need to spawn a child process which run in the backgrond and we keep taking input 
	// fmt.Println(cmd ,": needs to be exexuted ")
	// fmt.Println("args are :" , strings.Join(args," "))

	// first we need to check if a valid command or not 
	_ , err := exec.LookPath(command)
	if err != nil {
		fmt.Println(command + ": not found")
	}else {
		cmd := exec.Command(command, args[0:]...)
		err = cmd.Start()
		if err != nil { 
			log.Fatal(err)
		}
		// this is the process id now i need to wait for the process to finish and then decrement the counter of the job running 
		pid := cmd.Process.Pid
		fmt.Printf("[1] %d\n", pid)
		// procPath , err := getLinuxProcessPath(pid)
		// if err != nil {
		// 	fmt.Println(err)
		// }
		// fmt.Println(procPath)
	}
}
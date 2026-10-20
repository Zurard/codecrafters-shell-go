
package main

// this are the lest of commands we can execute in our code 
var toExit = false

var commands = map[string] struct{}{
	"exit" : struct{}{},
	"echo" : struct{}{},
}


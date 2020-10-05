package main

import (
	"fmt"
)

/*
HelloWorld a personlized approach!
*/
func HelloWorld() string {
	var name string
	fmt.Println("What is your name?")
	fmt.Scanln(&name)
	results := ("Hello, " + name)
	fmt.Printf(results)
	return results
}

func main() {
	HelloWorld()
}

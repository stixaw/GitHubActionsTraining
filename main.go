package main

import (
	"fmt"
)

/*
HelloWorld a personlized approach!
*/
func HelloWorld() {
	var name string
	fmt.Println("What is your name?")
	fmt.Scanln(&name)

	fmt.Printf("Hello, %s", name)
}

func main() {
	HelloWorld()
}

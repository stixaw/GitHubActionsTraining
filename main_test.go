package main_test

import (
	"fmt"
	"testing"
)

func helloWorld(t *testing.T) {
	name := "TIM"
	result := main.HelloWorld()
	fmt.Println(result)

}

package main

import "fmt"

// Hello says hello to the argument
func Hello(name string) string {
	if name == "" {
		return "Hello, World"
	} else {
		return "Hello, " + name
	}
}

func main() {
	fmt.Println(Hello("world"))
}

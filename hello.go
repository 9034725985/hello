package main

import "fmt"

const englishPrefix = "Hello, "

// Hello returns a string
func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishPrefix + name
}

func main() {
	fmt.Println(Hello("world"))
}

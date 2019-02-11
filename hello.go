package main

import "fmt"

const englishPrefix = "Hello, "
const spanishPrefix = "Hola, "
const spanish = "Spanish"

// Hello returns a string
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	if language == spanish {
		return spanishPrefix + name
	}
	return englishPrefix + name
}

func main() {
	fmt.Println(Hello("world", "English"))
}

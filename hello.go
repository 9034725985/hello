package main

import "fmt"

// Hello returns a string
func Hello(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}

func main() {
	fmt.Println(Hello("world"))
}

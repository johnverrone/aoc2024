package main

import "fmt"

func main() {
	b := Bar()
	fmt.Println("Hello, World!", b)
}

func Bar() string {
	return "Bar"
}

package main

import "fmt"

func main() {
	Hello("hello world")
}

func Hello(s string) string {
	fmt.Println(s)
	return s
}

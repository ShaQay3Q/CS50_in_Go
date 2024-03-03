package main

import (
	"fmt"
)

func main() {
	str := getLine("Write something: ")
	fmt.Printf("%s\n", str)
}

func get_string() string {
	var str string
	str = ""
	return str
}

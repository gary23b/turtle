package main

import (
	"encoding/base64"
	"fmt"
	"os"
)

func main() {
	fileData, err := os.ReadFile("turtle.png")
	if err != nil {
		panic(err)
	}
	result := base64.StdEncoding.EncodeToString(fileData)
	fmt.Printf("var turtleImage string = \"%s\"\n", result)

	fileData, err = os.ReadFile("arrow.png")
	if err != nil {
		panic(err)
	}
	result = base64.StdEncoding.EncodeToString(fileData)
	fmt.Printf("var arrowImage string = \"%s\"\n", result)
}

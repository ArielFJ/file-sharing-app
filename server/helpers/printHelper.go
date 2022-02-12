package helpers

import "fmt"

func PrintErr(err error) {
	fmt.Printf("Error: %v\n", err)
}

func Notify(message string) {
	fmt.Println("->", message)
}
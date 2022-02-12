package helpers

import "fmt"

func PrintErr(err error) {
	fmt.Printf("Error: %v", err)
}

func PrintErrPrefix(prefix string, err error) {
	fmt.Printf("Error %v: %v", prefix, err)
}

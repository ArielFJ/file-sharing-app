package main

import "fmt"

var (
	USERNAME = "username"
	CHANNEL  = "channel"
	SEND     = "send"
	MESSAGE  = "msg"
	LIST     = "list"
	EXIT     = "exit"
	HELP     = "help"
)

func expectResponse(cmd string) bool {
	switch cmd {
	case USERNAME, CHANNEL, SEND, LIST, EXIT:
		return true
	default:
		return false
	}
}

func showHelp() {
	fmt.Println("\tAvailable Commands:")
	fmt.Println("\t", HELP, ": Shows this help.")

	fmt.Println("\t", USERNAME, ": Set your username in the server.")

	fmt.Println("\t", CHANNEL, ": Connect or creates a channel.")
	fmt.Println("\t\tUses:", CHANNEL, "<channel # or name>")

	fmt.Println("\t", SEND, ": Send file to channel.")
	fmt.Println("\t\tUses:", SEND, "<channel # or name> <file path>")

	fmt.Println("\t", MESSAGE, ": Send message to channel.")
	fmt.Println("\t\tUses:", MESSAGE, "<channel # or name> <message>")

	fmt.Println("\t", LIST, ": List available channels.")

	fmt.Println("\t", EXIT, ": Close connection with the server.")
}

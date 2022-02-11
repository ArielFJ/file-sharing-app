package main

import (
	"fmt"
	"strings"
)

var (
	USERNAME = "username"
	CHANNEL  = "channel"
	STOP     = "stop"
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

func validateCommand(req request) (bool, string) {
	payload := strings.TrimSpace(req.Payload)
	args := strings.Split(payload, " ")

	switch req.Command {
	case USERNAME:
	case CHANNEL:
		if len(args) != 1 || payload == "" {
			return false, req.Command + " take one argument"
		}
	case SEND:

	case MESSAGE:

	case LIST:

	case EXIT:
	default:
		return false, "Invalid Command"
	}
	return true, ""
}

func showHelp() {
	fmt.Println("\tAvailable Commands:")
	fmt.Println("\t", HELP, ": Shows this help.")

	fmt.Println("\t", USERNAME, ": Gets or sets your username in the server.")
	fmt.Println("\t\tUses:", USERNAME, "<desired username, omit if you want to get your username>")

	fmt.Println("\t", CHANNEL, ": Connect or creates a channel.")
	fmt.Println("\t\tUses:", CHANNEL, "<channel # or name>")

	fmt.Println("\t", STOP, ": Closes a connection with a channel. You have to close the channel to continue using the tool normally.")

	fmt.Println("\t", SEND, ": Send file to channel.")
	fmt.Println("\t\tUses:", SEND, "<channel # or name> <file path>")

	fmt.Println("\t", MESSAGE, ": Send message to channel.")
	fmt.Println("\t\tUses:", MESSAGE, "<channel # or name> <message>")

	fmt.Println("\t", LIST, ": List available channels.")

	fmt.Println("\t", EXIT, ": Close connection with the server.")
}

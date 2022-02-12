package models

import (
	"fmt"
	"strings"
)

var (
	USERNAME = "username" // Get or set username
	CHANNEL  = "channel"  // Join a channel
	STOP     = "stop"     // Close the connection with a channel. While in a channel, this is the only valid command
	SEND     = "send"     // Send a file to a channel
	MESSAGE  = "msg"      // Send a message to a channel
	LIST     = "list"     // List available channels
	EXIT     = "exit"     // Close the app
	HELP     = "help"     // Show the help text
)

func validateCommand(req request) (bool, string) {
	payload := strings.TrimSpace(req.Payload)
	args := strings.Split(payload, " ")

	switch req.Command {
	case USERNAME:
		if len(args) > 1 {
			return false, req.Command + " takes zero or one argument. One for setting the username."
		}
	case CHANNEL:
		if len(args) != 1 || payload == "" {
			return false, req.Command + " takes one argument"
		}
	case SEND, MESSAGE:
		chanName := strings.TrimSpace(args[0])
		realPayload := strings.TrimSpace(strings.Join(args[1:], " "))
		if len(chanName) == 0 || len(realPayload) == 0 {
			return false, req.Command + " takes two arguments. The channel # or name, and the data to send."
		}
	case LIST, EXIT:
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

	fmt.Println("\t", STOP, ": Closes a connection with a channel.\n You have to close the channel to continue using the tool normally.")

	fmt.Println("\t", SEND, ": Send file to channel.")
	fmt.Println("\t\tUses:", SEND, "<channel # or name> <file path>")

	fmt.Println("\t", MESSAGE, ": Send message to channel.")
	fmt.Println("\t\tUses:", MESSAGE, "<channel # or name> <message>")

	fmt.Println("\t", LIST, ": List available channels.")

	fmt.Println("\t", EXIT, ": Close connection with the server.")
}

package handlers

import (
	"file-sharing-app/server/helpers"
	"file-sharing-app/server/models"
	"fmt"
)

// HandleUsername handles requests to change or return the username
func HandleUsername(c *models.Client, r models.Request) {
	oldUsername := c.Username
	response := models.NewResponse(models.OK, r.Command, fmt.Sprintf("Current username: %v\n", c.Username))
	if len(r.Payload) > 0 {
		c.Username = r.Payload
		response.Result = fmt.Sprintf("New username: %v\n", c.Username)
		helpers.Notify(fmt.Sprintf("%v change its username from %q", c.GetIdentifier(), oldUsername))
	}
	c.Send(response.ToBuffer())
}

// HandleExit handle the disconnection of the client
func HandleExit(c *models.Client, clients map[models.Client]bool) {
	helpers.Notify("Disconnecting user " + c.GetIdentifier())
	c.Disconnect()
	delete(clients, *c)
}
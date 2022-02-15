package handlers

import (
	"encoding/json"
	"file-sharing-app/server/models"
)

func HandleGetData(c *models.Client, channels map[string]*models.Channel) {
	channelsData := []models.ChannelResponse{}

	for _, channel := range channels {
		res := models.NewChannelResponse(*channel)
		channelsData = append(channelsData, *res)
	}

	jsonBytes, err := json.Marshal(channelsData)
	if err != nil {
		c.Send([]byte("Error: " + err.Error()))
	}

	c.Send(jsonBytes)
	c.Disconnect()
}

func HandleGetClients(c *models.Client, clients map[*models.Client]bool) {
	clientsData := []string{}

	for client := range clients {
		clientsData = append(clientsData, client.GetIdentifier())
	}

	jsonBytes, err := json.Marshal(clientsData)
	if err != nil {
		c.Send([]byte("Error: " + err.Error()))
	}

	c.Send(jsonBytes)
	c.Disconnect()
}
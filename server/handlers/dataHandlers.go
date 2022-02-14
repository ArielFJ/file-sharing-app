package handlers

import (
	"encoding/json"
	"file-sharing-app/server/models"
)

func HandleGetData(c *models.Client, r models.Request, channels map[string]*models.Channel) {
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
}
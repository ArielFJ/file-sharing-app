package handlers

import (
	"file-sharing-app/server/helpers"
	"file-sharing-app/server/models"
	"fmt"
	"path/filepath"
	"strings"
)

// HandleJoinChannel handles requests to join a channel. If the channel doesn't exists, it will be created.
func HandleJoinChannel(c *models.Client, r models.Request, channels map[string][]*models.Client) {
	chanName := r.Payload
	if len(chanName) < 1 {
		res := models.NewResponse(models.ERROR, r.Command, "Channel must have an identifier")
		c.Send(res.ToBuffer())
		return
	}
	clientsInChannel, exists := channels[chanName]
	var clients []*models.Client
	if !exists {
		clients = []*models.Client{c}
	} else {
		clients = append(clientsInChannel, c)
	}

	channels[chanName] = clients
	c.CurrentChannel = chanName

	helpers.Notify(fmt.Sprintf("%v connected to channel %q", c.GetIdentifier(), c.CurrentChannel))

	res := models.NewResponse(models.OK, r.Command, fmt.Sprintf("%v added to channel %q\n", c.Username, chanName))
	c.Send(res.ToBuffer())
}

// HandleQuitChannel handles disconnection from channels. If the channel is left with 0 clients, it will close.
func HandleQuitChannel(c *models.Client, r models.Request, channels map[string][]*models.Client) {
	res := models.NewResponse(models.ERROR, r.Command, fmt.Sprintf("%v removed from channel %q\n", c.Username, c.CurrentChannel))
	currentClients, channelExists := channels[c.CurrentChannel]
	if channelExists {
		// Remove the client from the channel
		clients := []*models.Client{}
		for _, client := range currentClients {
			if client != c {
				clients = append(clients, client)
			}
		}
		channels[c.CurrentChannel] = clients
	}

	// Close the channel if it doesn't have users
	if len(channels[c.CurrentChannel]) == 0 {
		delete(channels, c.CurrentChannel)
	}

	helpers.Notify(fmt.Sprintf("%v left channel %q", c.GetIdentifier(), c.CurrentChannel))
	c.CurrentChannel = ""

	c.Send(res.ToBuffer())
}

// HandleListChannels handles request to show a list of available channels.
func HandleListChannels(c *models.Client, r models.Request, channels map[string][]*models.Client) {
	channelsText := "Available Channels:\n"
	for chanName, clients := range channels {
		channelsText += fmt.Sprintf("\t - %v (%v clients)\n", chanName, len(clients))
	}

	if len(channels) == 0 {
		channelsText = "No Available Channels. Run help to see how to create one."
	}

	res := models.NewResponse(models.OK, r.Command, channelsText)
	c.Send(res.ToBuffer())
}

// HandleMessageToChannel handles requests to send messages to clients in a channel.
func HandleMessageToChannel(c *models.Client, r models.Request, channels map[string][]*models.Client) {
	chanName, realPayload := getChannelPayloadArgs(r.Payload)

	clients, err := checkChannel(channels, chanName, r.Command, c)
	if err != nil {
		return
	}

	broadcastDataToClients(c, clients, func() models.Response {
		res := models.NewResponse(models.OK, r.Command, fmt.Sprintf("MSG from %v: %v", c.Username, realPayload))
		return res
	})

	senderResponse := models.NewResponse(models.OK, r.Command, "Message sent to channel "+chanName)
	c.Send(senderResponse.ToBuffer())
}

// HandleSendFileToChannel handles requests to send files to clients in a channel.
func HandleSendFileToChannel(c *models.Client, r models.Request, channels map[string][]*models.Client) {
	chanName, filePath := getChannelPayloadArgs(r.Payload)
	fileName := filepath.Base(filePath)

	clients, err := checkChannel(channels, chanName, r.Command, c)
	if err != nil {
		return
	}

	broadcastDataToClients(c, clients, func() models.Response {
		res := models.NewResponse(models.OK, r.Command, fmt.Sprintf("File from %v: %v", c.Username, fileName))
		res.Data = r.Data
		return res
	})

	senderResponse := models.NewResponse(models.OK, r.Command, fmt.Sprintf("File %v sent to channel %v", fileName, chanName))
	c.Send(senderResponse.ToBuffer())
}

func checkChannel(channels map[string][]*models.Client, chanName, command string, c *models.Client) (clients []*models.Client, err error) {
	clients, channelExists := channels[chanName]
	if !channelExists {
		res := models.NewResponse(models.ERROR, command, fmt.Sprintf("Channel %v does not exists", chanName))
		c.Send(res.ToBuffer())
		return nil, fmt.Errorf("channel does not exists")
	}

	return clients, nil
}

func broadcastDataToClients(sender *models.Client, recipients []*models.Client, getResponse func() models.Response) {
	for _, client := range recipients {
		if client != sender {
			res := getResponse()
			client.Send(res.ToBuffer())
		}
	}
}

func getChannelPayloadArgs(payload string) (chanName, cmdArgs string) {
	args := strings.Split(strings.TrimSpace(payload), " ")
	chanName = strings.TrimSpace(args[0])
	cmdArgs = strings.TrimSpace(strings.Join(args[1:], " "))
	return
}

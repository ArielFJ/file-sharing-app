package handlers

import (
	"file-sharing-app/server/helpers"
	"file-sharing-app/server/models"
	"fmt"
	"path/filepath"
	"strings"
)

// HandleJoinChannel handles requests to join a channel. If the channel doesn't exists, it will be created.
func HandleJoinChannel(c *models.Client, r models.Request, channels map[string]*models.Channel) {
	chanName := r.Payload
	if len(chanName) < 1 {
		res := models.NewResponse(models.ERROR, r.Command, "Channel must have an identifier")
		c.Send(res.ToBuffer())
		return
	}

	channel, exists := channels[chanName]
	if !exists {
		channels[chanName] = models.NewChannel(chanName)
		channel = channels[chanName]
	}

	channel.ClientsJoined =  append(channel.ClientsJoined, c)

	c.CurrentChannel = chanName

	helpers.Notify(fmt.Sprintf("%v connected to channel %q", c.GetIdentifier(), c.CurrentChannel))

	res := models.NewResponse(models.OK, r.Command, fmt.Sprintf("%v added to channel %q\n", c.Username, chanName))
	c.Send(res.ToBuffer())
}

// HandleQuitChannel handles disconnection from channels. If the channel is left with 0 clients, it will close.
func HandleQuitChannel(c *models.Client, r models.Request, channels map[string]*models.Channel) {
	res := models.NewResponse(models.ERROR, r.Command, fmt.Sprintf("%v removed from channel %q\n", c.Username, c.CurrentChannel))
	channel, channelExists := channels[c.CurrentChannel]
	if channelExists {
		// Remove the client from the channel
		clients := []*models.Client{}
		for _, client := range channel.ClientsJoined {
			if client != c {
				clients = append(clients, client)
			}
		}
		channel.ClientsJoined = clients
	}

	// Close the channel if it doesn't have users
	// if len(channels[c.CurrentChannel]) == 0 {
	// 	delete(channels, c.CurrentChannel)
	// }

	helpers.Notify(fmt.Sprintf("%v left channel %q", c.GetIdentifier(), c.CurrentChannel))
	c.CurrentChannel = ""

	c.Send(res.ToBuffer())
}

// HandleListChannels handles request to show a list of available channels.
func HandleListChannels(c *models.Client, r models.Request, channels map[string]*models.Channel) {
	channelsText := "Available Channels:\n"
	for chanName, channel := range channels {
		channelsText += fmt.Sprintf("\t - %v (%v clients)\n", chanName, len(channel.ClientsJoined))
	}

	if len(channels) == 0 {
		channelsText = "No Available Channels. Run help to see how to create one."
	}

	res := models.NewResponse(models.OK, r.Command, channelsText)
	c.Send(res.ToBuffer())
}

// HandleMessageToChannel handles requests to send messages to clients in a channel.
func HandleMessageToChannel(c *models.Client, r models.Request, channels map[string]*models.Channel) {
	chanName, realPayload := getChannelPayloadArgs(r.Payload)

	channel, err := checkChannel(channels, chanName, r.Command, c)
	if err != nil {
		return
	}

	channel.AddMessage(fmt.Sprintf("%v: %v", c.GetIdentifier(), realPayload))

	broadcastDataToClients(c, channel.ClientsJoined, func() models.Response {
		res := models.NewResponse(models.OK, r.Command, fmt.Sprintf("MSG from %v: %v", c.Username, realPayload))
		return res
	})

	senderResponse := models.NewResponse(models.OK, r.Command, "Message sent to channel "+chanName)
	c.Send(senderResponse.ToBuffer())
}

// HandleSendFileToChannel handles requests to send files to clients in a channel.
func HandleSendFileToChannel(c *models.Client, r models.Request, channels map[string]*models.Channel) {
	chanName, filePath := getChannelPayloadArgs(r.Payload)
	fileName := filepath.Base(filePath)

	channel, err := checkChannel(channels, chanName, r.Command, c)
	if err != nil {
		return
	}

	channel.AddFile(fmt.Sprintf("%v sent: %v", c.GetIdentifier(), fileName), int64(len(r.Data)))

	broadcastDataToClients(c, channel.ClientsJoined, func() models.Response {
		res := models.NewResponse(models.OK, r.Command, fmt.Sprintf("File from %v: %v", c.Username, fileName))
		res.Data = r.Data
		return res
	})

	senderResponse := models.NewResponse(models.OK, r.Command, fmt.Sprintf("File %v sent to channel %v", fileName, chanName))
	c.Send(senderResponse.ToBuffer())
}

func checkChannel(channels map[string]*models.Channel, chanName, command string, c *models.Client) (channel *models.Channel, err error) {
	channel, channelExists := channels[chanName]
	
	if !channelExists {
		res := models.NewResponse(models.ERROR, command, fmt.Sprintf("Channel %v does not exists", chanName))
		c.Send(res.ToBuffer())
		return nil, fmt.Errorf("channel does not exists")
	}

	return channel, nil
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

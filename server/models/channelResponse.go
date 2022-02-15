package models

type ChannelResponse struct {
	Name           string   `json:"name"`
	TotalBytesSent int64    `json:"bytesSent"`
	Clients        []string `json:"clients"`
	FilesSent      []string `json:"filesSent"`
	MessagesSent   []string      `json:"messagesSent"`
}

func NewChannelResponse(c Channel) *ChannelResponse {
	res := &ChannelResponse{
		Name:           c.Name,
		TotalBytesSent: c.TotalBytesSent,
		Clients:        []string{},
		FilesSent:      c.FilesSent,
		MessagesSent:   c.MessagesSent,
	}

	for _, client := range c.ClientsJoined {
		res.Clients = append(res.Clients, client.GetIdentifier())
	}

	return res
}

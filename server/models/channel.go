package models

type Channel struct {
	Name           string    `json:"name"`
	TotalBytesSent int64     `json:"bytesSent"`
	ClientsJoined  []*Client `json:"clientsJoined"`
	FilesSent      []string  `json:"filesSent"`
	MessagesSent   []string  `json:"messagesSent"`
}

func NewChannel(n string) *Channel {
	return &Channel{
		Name:           n,
		TotalBytesSent: 0,
		ClientsJoined:  []*Client{},
		FilesSent:      []string{},
		MessagesSent:   []string{},
	}
}

func (c *Channel) AddFile(fileName string, fileSize int64) {
	c.FilesSent = append(c.FilesSent, fileName)
	c.TotalBytesSent += fileSize
}

func (c *Channel) AddMessage(message string) {
	c.MessagesSent = append(c.MessagesSent, message)
}

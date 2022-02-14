package models

type ChannelData struct {
	Name           string   `json:"name"`
	TotalBytesSent int64    `json:"bytesSent"`
	Clients        []string `json:"clients"`
	FilesSent      []string `json:"filesSent"`
	MessagesSent   []string `json:"messagesSent"`
}

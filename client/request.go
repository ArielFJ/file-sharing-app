package main

import (
	"file-sharing-app/client/helpers"
	"fmt"
	"strings"
)

type request struct {
	Payload string `json:"payload"`
	Data     []byte `json:"data"`
	Command  string `json:"command"`
}

func NewRequest(cmd, payload string) request {
	return request{
		Command: cmd,
		Data: []byte{},
		Payload: payload,
	}
}

func BuildRequest(text string) (bool, request) {
	cleanText := strings.ReplaceAll(text, "\r\n", "") // Take just the input without the return
	words := strings.Split(cleanText, " ")
	cmd := strings.TrimSpace(helpers.NormalizeString(words[0]))
	payload := strings.TrimSpace(strings.Join(words[1:], " "))

	if len(cmd) == 0{
		return false, request{}
	}

	return true, NewRequest(cmd, payload)
}

func (m *request) String() string {
	return fmt.Sprintf("{\ncommand: %v,\n payload: %v,\n dataLength: %v \n}\n", m.Command, m.Payload, len(m.Data))
}

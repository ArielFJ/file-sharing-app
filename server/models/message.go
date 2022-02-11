package models

import "fmt"

type message struct {
	Payload string `json:"payload"`
	Data     []byte `json:"data"`
	Command  string `json:"command"`
}

func NewMessage(cmd, payload string) message {
	return message{
		Command: cmd,
		Data: []byte{},
		Payload: payload,
	}
}

func (m *message) String() string {
	return fmt.Sprintf("{\ncommand: %v,\n payload: %v,\n dataLength: %v \n}\n", m.Command, m.Payload, len(m.Data))
}

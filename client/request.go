package main

import "fmt"

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

func (m *request) String() string {
	return fmt.Sprintf("{\ncommand: %v,\n payload: %v,\n dataLength: %v \n}\n", m.Command, m.Payload, len(m.Data))
}

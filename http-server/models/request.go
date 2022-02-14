package models

import "fmt"

type Request struct {
	Payload string `json:"payload"`
	Data     []byte `json:"data"`
	Command  string `json:"command"`
}

func NewRequest(cmd, payload string) Request {
	return Request{
		Command: cmd,
		Data: []byte{},
		Payload: payload,
	}
}

func (m *Request) String() string {
	return fmt.Sprintf("{\ncommand: %v,\n payload: %v,\n dataLength: %v \n}\n", m.Command, m.Payload, len(m.Data))
}

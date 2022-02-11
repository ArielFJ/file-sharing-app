package models

import "fmt"

type message struct {
	Filename string `json:"filename"`
	Data     []byte `json:"data"`
	Command  string `json:"command"`
}

func (m *message) String() string {
	return fmt.Sprintf("{\ncommand: %v,\n filename: %v,\n dataLength: %v \n}\n", m.Command, m.Filename, len(m.Data))
}

package models

import (
	"file-sharing-app/client/helpers"
	"fmt"
	"os"
	"strings"
)

type request struct {
	Payload string `json:"payload"`
	Data    []byte `json:"data"`
	Command string `json:"command"`
}

func NewRequest(cmd, payload string) request {
	return request{
		Command: cmd,
		Data:    []byte{},
		Payload: payload,
	}
}

func BuildRequest(text string) (bool, request) {
	cleanText := strings.ReplaceAll(text, "\r\n", "") // Take just the input without the return
	words := strings.Split(cleanText, " ")
	cmd := strings.TrimSpace(helpers.NormalizeString(words[0]))
	payload := strings.TrimSpace(strings.Join(words[1:], " "))

	if len(cmd) == 0 {
		return false, request{}
	}

	req := NewRequest(cmd, payload)

	return true, req
}

func (r *request) String() string {
	return fmt.Sprintf("{\ncommand: %v,\n payload: %v,\n dataLength: %v \n}\n", r.Command, r.Payload, len(r.Data))
}

func (r *request) UpdateRequestWithFileData() error {
	args := strings.Split(strings.TrimSpace(r.Payload), " ")
	filePath := strings.TrimSpace(strings.Join(args[1:], " "))

	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		helpers.PrintErrPrefix("FILE", err)
		return err
	}

	r.Data = fileContent
	return nil
}

package models

import (
	"encoding/json"
	"fmt"
)

const (
	OK int = iota
	ERROR
)

type Response struct {
	Code    int    `json:"code"`
	Command string `json:"command"`
	Result  string `json:"result"`
	Data    []byte `json:"data"`
}

func NewResponse(code int, command, result string) Response {
	return Response{
		Code:    code,
		Command: command,
		Result:  result,
		Data:    []byte{},
	}
}

func NewResponseWithDefaultCode(code int, result string) Response {
	return NewResponse(code, "", result)
}

func (r *Response) String() string {
	return fmt.Sprintf("{\n code: %v,\n command: %v,\n result: %v,\n dataLength: %v \n}\n", r.Code, r.Command, r.Result, len(r.Data))
}

func (r *Response) ToBuffer() []byte {
	bytes, err := json.Marshal(r)
	if err != nil {
		return []byte{}
	}
	return bytes
}
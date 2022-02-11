package main

import (
	"encoding/json"
	"fmt"
)

const (
	OK int = iota
	ERROR
)

type response struct {
	Code   int    `json:"code"`
	Result string `json:"result"`
	Data   []byte `json:"data"`
}

func NewResponse(code int, result string) response {
	return response{
		Code: code,
		Result: result,
		Data:    []byte{},
	}
}

func (r *response) String() string {
	return fmt.Sprintf("{\n code: %v,\n result: %v,\n dataLength: %v \n}\n", r.Code, r.Result, len(r.Data))
}

func (r *response) ToBuffer() []byte {
	bytes, err := json.Marshal(r)
	if err != nil {
		return []byte{}
	}
	return bytes
}
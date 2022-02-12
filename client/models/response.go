package models

import (
	"bufio"
	"encoding/json"
	"file-sharing-app/client/helpers"
	"fmt"
	"io"
	"net"
)

const (
	OK int = iota
	ERROR
)

type Response struct {
	Code   int    `json:"code"`
	Result string `json:"result"`
	Data   []byte `json:"data"`
}

func NewResponse(code int, result string) Response {
	return Response{
		Code: code,
		Result: result,
		Data:    []byte{},
	}
}

func (r *Response) String() string {
	return fmt.Sprintf("{\n code: %v,\n result: %v,\n dataLength: %v \n}\n", r.Code, r.Result, len(r.Data))
}

func (r *Response) ToBuffer() []byte {
	bytes, err := json.Marshal(r)
	if err != nil {
		return []byte{}
	}
	return bytes
}

func ReadFromConn(c net.Conn) (res Response, mildError, fatalError bool) {
	fatalError = false
	mildError = false

	netData, err := bufio.NewReader(c).ReadString('\n')
	if err != nil {
		helpers.PrintErrPrefix("RESPONSE", err)
		if err == io.EOF {
			fatalError = true
		}

		mildError = true
	}

	err = json.Unmarshal([]byte(netData), &res)
	if err != nil {
		res = NewResponse(ERROR, "")
	}

	return res, mildError, fatalError
}
package api

import (
	"encoding/json"
)

type Response struct {
	Success    bool            `json:"success"`
	ServerTime int             `json:"server_time"`
	Meta       string          `json:"meta"`
	Message    json.RawMessage `json:"message"`
	Error      string          `json:"error"`
	Data       json.RawMessage `json:"data,omitempty"`
}

func ResponseParse(resp []byte) (*Response, error) {
	response := &Response{}
	err := json.Unmarshal(resp, response)
	return response, err
}

package api

import "encoding/json"

type ResponseError struct {
	message string
}

func (re *ResponseError) Error() string {
	return re.message
}

type Response struct {
	Success    bool            `json:"success"`
	ServerTime int             `json:"server_time"`
	Meta       string          `json:"meta"`
	Message    string          `json:"message"`
	Error      string          `json:"error"`
	Data       json.RawMessage `json:"data"`
}

func ResponseParse(resp []byte) (*Response, error) {
	response := &Response{}
	err := json.Unmarshal(resp, response)
	return response, err
}

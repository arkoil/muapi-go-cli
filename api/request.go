package api

import (
	"crypto/sha256"
	"encoding/hex"
	"net/url"
	"strings"
)

type Request struct {
	sign      string `json:"sign"`
	publicKey string `json:"public_key"`
	data      string `json:"data"`
	meta      string
}

type RequestQuery map[string]string

func (api MUAPI) Sign(data string) string {
	h := sha256.New()
	h.Write([]byte(data + api.privateKey))
	sum := hex.EncodeToString(h.Sum(nil))
	return strings.ToLower(sum)
}

func (api MUAPI) NewRequest(data string, meta string) Request {
	sign := api.Sign(data)
	return Request{sign, api.publicKey, data, meta}
}

func (req Request) ToVal() (val url.Values) {
	val = url.Values{}
	val.Add("sign", req.sign)
	val.Add("public_key", req.publicKey)
	val.Add("data", req.data)
	return
}

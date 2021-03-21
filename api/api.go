package api

import (
	"crypto/sha256"
	"net/http"
	"net/url"
)

var Endpoints = map[string]string{
	"auth": "/auth",
}

type MUAPI struct {
	client *http.Client
	privateKey string
	publicKey string
	host string
	port string
	Resource Resource
}

func New(client *http.Client, privKey string, pubKey string, resName string,host string, port string) *MUAPI {
	res := Resource{resName}
	return &MUAPI{client: client, privateKey: privKey, publicKey: pubKey, Resource: res, host: host, port: port}
}
func (api MUAPI) Sign(data string) string {
	h := sha256.New()
	h.Write([]byte(data + api.privateKey))
	sum := h.Sum(nil)
	return string(sum)
}

func (api MUAPI) Request2Api(endpoint string) (string, error){
	uri := api.host + api.port + Endpoints[endpoint]
	data := url.Values{}
	res, err := api.client.PostForm(uri, data)
	return res.Status, err
}
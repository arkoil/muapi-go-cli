package muapi_go_cli

import (
	"github.com/gpbbit/muapi-go-cli/api"
	"net/http"
)

func NewAPI(client *http.Client, privKey string, pubKey string, resName string, resUrl string, host string, port string) (*api.MUAPI, error) {
	apiCli, err := api.New(client, privKey, pubKey, resName, resUrl, host, port)
	return apiCli, err
}

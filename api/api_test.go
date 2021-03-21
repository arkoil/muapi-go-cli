package api

import (
	"net/http"
	"testing"
)

func TestNewApi(t *testing.T) {
	client := &http.Client{}
	api := New(client, "qqqq", "xxxx", "wb", "http://localhost", "4000")
	defVal := "wb"
	if defVal != api.Resource.Name {
		t.Error("Bad resource name!")
	}
}

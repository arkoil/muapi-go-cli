package api

import (
	"net/http"
	"testing"
)

const ValidResourceName = "wb"
const ValidResourceURL = "https://www.wildberries.ru/"

func TestNewApi(t *testing.T) {
	client := &http.Client{}
	api, err := New(client, "qqqq", "xxxx", ValidResourceName, ValidResourceURL, "http://localhost", ":4000")
	defVal := "wb"
	if err != nil {
		t.Error(err)
	}

	if defVal != api.Resource.Name {
		t.Error("Bad resource name!")
	}
}

func TestMUAPI_Auth(t *testing.T) {
	client := &http.Client{}
	api, err := New(client, "qqqq", "xxxx", ValidResourceName, ValidResourceURL, "http://localhost", ":4000")
	if err != nil {
		t.Error(err)
	}
	_, err = api.Auth()
	if err != nil {
		t.Error(err)
	}

}

func TestMUAPI_ResourceGet(t *testing.T) {
	client := &http.Client{}
	api, err := New(client, "qqqq", "xxxx", ValidResourceName, ValidResourceURL, "http://localhost", ":4000")
	if err != nil {
		t.Error(err)
	}
	_, err = api.ResourceGet()
	if err != nil {
		t.Error(err)
	}
}

func TestMUAPI_Catalogs(t *testing.T) {
	client := &http.Client{}
	api, err := New(client, "qqqq", "xxxx", ValidResourceName, ValidResourceURL, "http://localhost", ":4000")
	if err != nil {
		t.Error(err)
	}
	query := make(map[string]string)
	_, err = api.Catalogs(query)
	if err != nil {
		t.Error(err)
	}
}

func TestMUAPI_CatalogAdd(t *testing.T) {
	client := &http.Client{}
	api, err := New(client, "qqqq", "xxxx", ValidResourceName, ValidResourceURL, "http://localhost", ":4000")
	if err != nil {
		t.Error(err)
	}
	catalog := Catalog{Name: "одежда", URL: "https://www.wildberries.ru/catalog/muzhchinam/odezhda", Region: "Moscow", ParentId: "604e0c39c11b3a7dc7a35d01"}
	res, err := api.CatalogAdd(catalog)
	t.Log(res)
	if err != nil {
		t.Error(err)
	}
}

package muapi_go_cli

import (
	cli "github.com/gpbbit/muapi-go-cli/api"
	"net/http"
	"testing"
)

const ValidResourceName = "wb"
const ValidResourceURL = "https://www.wildberries.ru/"

func TestNewApi(t *testing.T) {
	client := &http.Client{}
	api, err := NewAPI(client, "qqqq", "xxxx", ValidResourceName, ValidResourceURL, "http://localhost", ":4000")
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
	api, err := NewAPI(client, "qqqq", "xxxx", ValidResourceName, ValidResourceURL, "http://localhost", ":4000")
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
	api, err := NewAPI(client, "qqqq", "xxxx", ValidResourceName, ValidResourceURL, "http://localhost", ":4000")
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
	api, err := NewAPI(client, "qqqq", "xxxx", ValidResourceName, ValidResourceURL, "http://localhost", ":4000")
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
	api, err := NewAPI(client, "qqqq", "xxxx", ValidResourceName, ValidResourceURL, "http://localhost", ":4000")
	if err != nil {
		t.Error(err)
	}
	catalog := cli.Catalog{
		Name: "брюки",
		URL: "https://www.wildberries.ru/catalog/muzhchinam/odezhda/bryuki-i-shorty",
		Region: "Moscow",
		ParentId: "604de597c11b3a65e50319f5",
	}
	_, err = api.CatalogAdd(catalog)
	if err != nil {
		t.Error(err)
	}
}

func TestMUAPI_Items(t *testing.T) {
	client := &http.Client{}
	api, err := NewAPI(client, "qqqq", "xxxx", ValidResourceName, ValidResourceURL, "http://localhost", ":4000")
	if err != nil {
		t.Error(err)
	}
	query := make(map[string]string)
	_, err = api.Items(query)
	if err != nil {
		t.Error(err)
	}
}

func TestMUAPI_ItemAdd(t *testing.T) {
	client := &http.Client{}
	api, err := NewAPI(client, "qqqq", "xxxx", ValidResourceName, ValidResourceURL, "http://localhost", ":4000")
	if err != nil {
		t.Error(err)
	}
	catalogs := make([]string, 1)
	catalogs[0] = "604de597c11b3a65e50319f5"
	item := cli.Item{
		Name:      "Футболка",
		URL:       "https://www.wildberries.ru/catalog/11566676/detail.aspx?targetUrl=GP",
		Region:    "Moscow",
		CatalogId: "605af44fc11b3a64ede0c961",
		Catalogs:  catalogs,
	}
	_, err = api.ItemAdd(item)
	if err != nil {
		t.Error(err)
	}
}

func TestMUAPI_CatalogFindByIds(t *testing.T) {
	client := &http.Client{}
	api, err := NewAPI(client, "qqqq", "xxxx", ValidResourceName, ValidResourceURL, "http://localhost", ":4000")
	if err != nil {
		t.Error(err)
	}
	catalogs := make([]string, 1)
	catalogs[0] = "604de597c11b3a65e50319f5"
	_, err = api.CatalogFindByIds(catalogs)
	if err != nil {
		t.Error(err)
	}
}

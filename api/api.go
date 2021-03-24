package api

import (
	"encoding/json"
	"io"
	"net/http"
)

// API routing list
var Endpoints = map[string]string{
	"auth":          "/auth",
	"resources":     "/resources",
	"resource":      "/resource",
	"resourceAdd":   "/resource/add",
	"catalogs":      "/resource/catalog",
	"catalogAdd":    "/resource/catalog/add",
	"items":         "/resource/catalog/item",
	"itemAdd":       "/resource/catalog/item/add",
	"getItemStruct": "/resource/get_item_struct",
	"searchItem":    "/resource/search/item",
	"searchCatalog": "/resource/search/catalog",
}

type MUAPI struct {
	Client     *http.Client // HTTP client
	privateKey string       // API private key
	publicKey  string       // API public key
	host       string       // API host address
	port       string       // API port at first :
	Resource   Resource
}

func New(client *http.Client, privKey string, pubKey string, resName string, resUrl string, host string, port string) (*MUAPI, error) {
	res := Resource{Name: resName, URL: resUrl}
	api := MUAPI{Client: client, privateKey: privKey, publicKey: pubKey, Resource: res, host: host, port: port}
	_, err := api.Auth()
	if err != nil {
		return &api, err
	}
	resourceJSON, err := api.ResourceAdd(res)
	resources := make([]Resource, 1)
	resources[0] = res
	err = json.Unmarshal([]byte(resourceJSON), &resources)
	if err != nil {
		return &api, err
	}
	api.Resource = resources[0]
	return &api, nil
}

func checkMeta(meta *string, paramMeta []string) {
	if len(paramMeta) > 0 {
		meta = &paramMeta[0]
	}

}

func (api MUAPI) Auth(paramMeta ...string) (string, error) {
	data := AuthData{Data: NewData()}
	var meta string
	checkMeta(&meta, paramMeta)
	_, res, err := api.Request2Api("auth", data, meta)
	if err != nil {
		return "", err
	}
	if !res.Success {
		err = &ResponseError{
			message: res.Error,
		}
	}
	return string(res.Data), err

}
func (api MUAPI) ResourceGet(paramMeta ...string) (string, error) {
	data := ResourceData{Resource: api.Resource.Name, Data: NewData()}
	var meta string
	checkMeta(&meta, paramMeta)
	_, res, err := api.Request2Api("resource", data, meta)
	if err != nil {
		return "", err
	}
	if !res.Success {
		err = &ResponseError{
			message: res.Error,
		}
	}
	return string(res.Data), err
}
func (api MUAPI) ResourceAdd(resource Resource, paramMeta ...string) (string, error) {
	data := ResourceAddData{Resource: resource, Data: NewData()}
	var meta string
	checkMeta(&meta, paramMeta)
	_, res, err := api.Request2Api("resourceAdd", data, meta)
	if err != nil {
		return "", err
	}
	if !res.Success {
		err = &ResponseError{
			message: res.Error,
		}
	}
	return string(res.Data), err

}
func (api MUAPI) Catalogs(catalogsQuery RequestQuery, paramMeta ...string) (string, error) {
	data := CatalogsData{ResourceName: api.Resource.Name, CatalogQuery: catalogsQuery, Data: NewData()}
	var meta string
	checkMeta(&meta, paramMeta)
	_, res, err := api.Request2Api("catalogs", data, meta)
	if err != nil {
		return "", err
	}
	if !res.Success {
		err = &ResponseError{
			message: res.Error,
		}
	}
	return string(res.Data), err
}
func (api MUAPI) CatalogAdd(catalog Catalog, paramMeta ...string) (string, error) {
	err := catalog.Validate()
	if err != nil {
		return "", err
	}
	data := CatalogAddData{ResourceName: api.Resource.Name, Catalog: catalog, Data: NewData()}
	//TODO check parent
	var meta string
	checkMeta(&meta, paramMeta)
	_, res, err := api.Request2Api("catalogAdd", data, meta)
	if err != nil {
		return "", err
	}
	if !res.Success {
		err = &ResponseError{
			message: res.Error,
		}
	}
	return string(res.Data), err

}
func (api MUAPI) Items(itemsQuery RequestQuery, paramMeta ...string) (string, error) {
	data := ItemsData{ResourceName: api.Resource.Name, ItemQuery: itemsQuery, Data: NewData()}
	var meta string
	checkMeta(&meta, paramMeta)
	_, res, err := api.Request2Api("items", data, meta)
	if err != nil {
		return "", err
	}
	if !res.Success {
		err = &ResponseError{
			message: res.Error,
		}
	}
	return string(res.Data), err
}
func (api MUAPI) ItemAdd(item Item, paramMeta ...string) (string, error) {
	data := ItemAddData{ResourceName: api.Resource.Name, Item: item, Data: NewData()}
	//TODO check catalog
	var meta string
	checkMeta(&meta, paramMeta)
	_, res, err := api.Request2Api("itemAdd", data, meta)
	if err != nil {
		return "", err
	}
	if !res.Success {
		err = &ResponseError{
			message: res.Error,
		}
	}
	return string(res.Data), err

}
func (api MUAPI) CatalogSearch(searchData CatalogSearchData, paramMeta ...string) (*Response, error) {
	var meta string
	checkMeta(&meta, paramMeta)
	_, res, err := api.Request2Api("searchCatalog", searchData, meta)
	return res, err

}
func (api MUAPI) CatalogFindByIds(ids []string, paramMeta ...string) (string, error) {
	searchData := CatalogSearchData{
		ResourceName: api.Resource.Name,
		SearchType:   "by_ids",
		IDField:      "_id",
		IDValues:     ids,
		Data:         NewData(),
		Request:      make(map[string]string),
	}
	var meta string
	checkMeta(&meta, paramMeta)
	res, err := api.CatalogSearch(searchData, meta)
	if err != nil {
		return "", err
	}
	if !res.Success {
		err = &ResponseError{
			message: res.Error,
		}
	}
	return string(res.Data), err

}
func (api MUAPI) CatalogCheckParents(ids []string) error {
	catalogs := make([]Catalog, len(ids))
	res, err := api.CatalogFindByIds(ids)
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(res), catalogs)
	if err != nil {
		return err
	}
	return nil
}
func (api MUAPI) Request2Api(endpoint string, data RequestData, meta string) (string, *Response, error) {
	uri := api.host + api.port + Endpoints[endpoint]
	strData, err := DataToFormat(data)
	errResp := &Response{Success: false}
	if err != nil {
		return "", errResp, err
	}
	request := api.NewRequest(strData, meta)
	res, err := api.Client.PostForm(uri, request.ToVal())
	if err != nil {
		return "", errResp, err
	}
	defer res.Body.Close()
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return "", errResp, err
	}
	response, err := ResponseParse(resBody)
	if err != nil {
		return "", errResp, err
	}
	return res.Status, response, err
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

package api

import (
	"encoding/json"
	"strconv"
	"time"
)

type RequestData interface {
	Info() (string, error)
}

//Base data
type Data struct {
	Time int64 `json:"time"`
}

//Struct data for auth request
type AuthData struct {
	Data
}

//Struct data for resource request
type ResourceData struct {
	Resource string `json:"resource"`
	Data
}

//Struct data for add resource request
type ResourceAddData struct {
	Resource Resource `json:"resource"`
	Data
}

//Struct data for catalogs request
type CatalogsData struct {
	ResourceName string            `json:"resource"`
	CatalogQuery map[string]string `json:"catalog"`
	Data
}

//Struct data for add catalog request
type CatalogAddData struct {
	ResourceName string  `json:"resource"`
	Catalog      Catalog `json:"catalog"`
	Data
}

//Struct data for items request
type ItemsData struct {
	ResourceName string            `json:"resource"`
	ItemQuery    map[string]string `json:"item"`
	Data
}

//Struct data for add item request
type ItemAddData struct {
	ResourceName string `json:"resource"`
	Item         Item   `json:"item"`
	Data
}

type CatalogSearchData struct {
	ResourceName string            `json:"resource"`
	SearchType   string            `json:"type"`
	IDField      string            `json:"id_field,omitempty"`
	IDValues     []string          `json:"id_values,omitempty"`
	Request      map[string]string `json:"request"`
	Data
}

//Creating Data for time now
func NewData() Data {
	now := time.Now()
	return Data{
		Time: now.Unix(),
	}
}

// Marshaling RequestData to JSON format
func DataToFormat(data RequestData) (string, error) {
	jData, err := json.Marshal(data)
	return string(jData), err
}
func (data Data) Info() (string, error) {
	return strconv.FormatInt(data.Time, 10), nil
}

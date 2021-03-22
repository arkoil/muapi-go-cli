package api

import (
	"encoding/json"
	"strconv"
	"time"
)

type RequestData interface {
	Info() (string, error)
}

type Data struct {
	Time int64 `json:"time"`
}

type AuthData struct {
	Data
}
type ResourceData struct {
	Resource string `json:"resource"`
	Data
}

type ResourceAddData struct {
	Resource Resource `json:"resource"`
	Data
}

type CatalogsData struct {
	ResourceName string            `json:"resource"`
	CatalogQuery map[string]string `json:"catalog"`
	Data
}
type CatalogAddData struct {
	ResourceName string  `json:"resource"`
	Catalog      Catalog `json:"catalog"`
	Data
}

func NewData() Data {
	now := time.Now()
	return Data{
		Time: now.Unix(),
	}
}

func DataToFormat(data RequestData) (string, error) {
	jData, err := json.Marshal(data)
	return string(jData), err
}
func (data Data) Info() (string, error) {
	return strconv.FormatInt(data.Time, 10), nil
}

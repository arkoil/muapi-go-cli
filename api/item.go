package api

import "encoding/json"

type Item struct {
	Name        string                     `json:"name"`
	Article     string                     `json:"article,omitempty"`
	Brand       string                     `json:"brand,omitempty"`
	Model       string                     `json:"model,omitempty"`
	Rating      float64                    `json:"rating,omitempty"`
	Reviews     int                        `json:"reviews,omitempty"`
	Region      string                     `json:"region"`
	Description string                     `json:"description,omitempty"`
	ExternalId  string                     `json:"external_id,omitempty"`
	Catalogs    []string                   `json:"catalogs,omitempty"`
	CatalogId   string                     `json:"catalog_id"`
	Image       string                     `json:"image,omitempty"`
	Price       float64                    `json:"price,omitempty"`
	Currency    string                     `json:"currency,omitempty"`
	URL         string                     `json:"url"`
	Additional  map[string]json.RawMessage `json:"additional,omitempty"`
	UniformName string                     `json:"uniform_name,omitempty"`
	ResourceId  string                     `json:"resource_id"`
	CreatedAt   string                     `json:"created_at,omitempty"`
	UpdatedAt   string                     `json:"updated_at,omitempty"`
	ID          string                     `json:"id,omitempty"`
}

func (item Item) Validate() error {
	var errors []string
	if item.Name == "" {
		errors = append(errors, "the 'name' must exist ")
	}
	if item.URL == "" {
		errors = append(errors, "the 'url' must exist ")
	}
	if item.Region == "" {
		errors = append(errors, "the 'name' must exist ")
	}
	if len(errors) > 0 {
		return &RequestValidationError{messages: errors, model: "Item", method: "Validate"}
	}
	return nil
}

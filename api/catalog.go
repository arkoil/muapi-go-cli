package api

type Catalog struct {
	Name        string            `json:"name"`
	Description string            `json:"description,omitempty"`
	ExternalId  string            `json:"external_id,omitempty"`
	Region      string            `json:"region"`
	Parents     []string          `json:"parents,omitempty"`
	ParentId    string            `json:"parent_id,omitempty"`
	Image       string            `json:"image,omitempty"`
	URL         string            `json:"url"`
	Additional  map[string]string `json:"additional,omitempty"`
	UniformName string            `json:"uniform_name,omitempty"`
	ResourceId  string            `json:"resource_id"`
	ID          string            `json:"id,omitempty"`
}

func (catalog Catalog) Validate() error {
	var errors []string
	if catalog.Name == "" {
		errors = append(errors, "the 'name' must exist ")
	}
	if catalog.URL == "" {
		errors = append(errors, "the 'url' must exist ")
	}
	if catalog.Region == "" {
		errors = append(errors, "the 'name' must exist ")
	}
	if len(errors)>0 {
		return &RequestValidationError{messages: errors, model: "Catalog", method: "Validate"}
	}
	return nil
}

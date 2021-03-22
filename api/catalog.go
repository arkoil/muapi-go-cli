package api

type Catalog struct {
	Name string `json:"name"`
	Description string `json:"description"`
	ExternalId string `json:"external_id"`
	Region string `json:"region"`
	Parents []string `json:"parents"`
	ParentId string `json:"parent_id"`
	Image string `json:"image"`
	URL string `json:"url"`
	Additional map[string]string `json:"additional"`
	UniformName string `json:"uniform_name"`
	ResourceId string `json:"resource_id"`
}

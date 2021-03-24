package api

type Resource struct {
	Name      string `json:"name"`
	URL       string `json:"url"`
	Id        string `json:"id,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
}

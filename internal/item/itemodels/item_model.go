package itemodels

type CreateNewItemRequest struct {
	Name     string `json:"name"`
	Image    string `json:"image"`
	Quantity int    `json:"quantity"`
}

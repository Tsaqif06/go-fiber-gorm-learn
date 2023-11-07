package request

type BookCreateRequest struct {
	Title  string `json:"title" validate:"required"`
	Author string `json:"author"`
	Cover  string `json:"cover"`
}

type BookUpdateRequest struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Cover  string `json:"cover"`
}
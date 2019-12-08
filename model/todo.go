package model

// ToDo entity model
type ToDo struct {
	ID    int64  `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
}

package requests

type Todo struct {
	Title       string  `json:"title"`
	Description *string `json:"description,omitempty"`
	Completed   bool    `json:"completed"`
}
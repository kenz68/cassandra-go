package model

// Todo is model using for database model and JSON response model
type Todo struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

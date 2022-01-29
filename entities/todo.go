package entities

type Todo struct {
	Title string `json:"title"`
	Description string `json:"-"`
	IsCompleted bool `json:"isCompleted"`
}
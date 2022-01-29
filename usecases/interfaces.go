package usecases

import "github.com/storyofhis/golang-react-todo/entities"

type TodosRepository interface {
	GetAllTodos() ([]entities.Todo, error)
}
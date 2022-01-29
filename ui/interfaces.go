package ui

import "github.com/storyofhis/golang-react-todo/entities"

type Service interface {
	GetAllTodos() ([]entities.Todo, error)
}
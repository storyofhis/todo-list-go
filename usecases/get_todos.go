package usecases

import (
	"github.com/storyofhis/golang-react-todo/entities"
)

func GetTodos(repo TodosRepository) ([]entities.Todo, error){
	// return nil, ErrInternal
	todos, err := repo.GetAllTodos()
	if err != nil {
		return nil, ErrInternal
	}else {
		return todos, nil
	}
}
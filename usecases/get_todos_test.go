package usecases_test

import (
	"fmt"
	"testing"

	"github.com/gomagedon/expectate"
	"github.com/storyofhis/golang-react-todo/entities"
	"github.com/storyofhis/golang-react-todo/usecases"
)

var dummyTodos = []entities.Todo{
	{
		Title: "todo 1",
		Description: "description of todo 1",
		IsCompleted: true,
	},
	{
		Title: "todo 2",
		Description: "description of todo 2",
		IsCompleted: false,
	},
	{
		Title: "todo 3",
		Description: "description of todo 3",
		IsCompleted: true,
	},
}

type BadTodosRepo struct {
}

func (BadTodosRepo) GetAllTodos() ([]entities.Todo, error) {
	return nil, fmt.Errorf("something went wrong when getting all")
}

type MockTodosRepo struct {
}

func (MockTodosRepo) GetAllTodos() ([]entities.Todo, error) {
	return dummyTodos, nil 
}

func TestGetTodos( t *testing.T){
	// TEST
	t.Run("Returns ErrInternal when TodosRepository returns Error", func (t *testing.T) {
		expect := expectate.Expect(t)
		repo := new(BadTodosRepo)
		todos, err := usecases.GetTodos(repo)

		expect(err).ToBe(usecases.ErrInternal)
		// if err != usecases.ErrInternal {
		// 	t.Fatalf("Expected Error Internal, got: %v", err)
		// }

		// expect(todos).ToBe(nil)
		if todos != nil {
			t.Fatalf("Expected todos, got: %v", todos)
		}
	})
	// TEST
	t.Run("Returns todos from TodosRepository", func (t *testing.T){
		expect := expectate.Expect(t)
		repo := new(MockTodosRepo)
		todos, err := usecases.GetTodos(repo)
		
		expect(err).ToBe(nil)
		expect(todos).ToEqual(dummyTodos)
	})
}
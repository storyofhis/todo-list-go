package ui_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gomagedon/expectate"
	"github.com/storyofhis/golang-react-todo/entities"
	"github.com/storyofhis/golang-react-todo/ui"
)

type MockService struct {
	err error
	todos []entities.Todo
}

func (s MockService) GetAllTodos() ([]entities.Todo, error){
	if s.err != nil {
		return nil, s.err
	}
	return s.todos, nil
}

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
// refactor 
type HTTPTest struct {
	name string

	service *MockService
	inputMethod string
	inputURL string

	expectedStatus int
	expectedTodos []entities.Todo
}

func TestHTTP (t *testing.T) {
	
	tests := []HTTPTest{
		{
			name: "Random Error gives 500 status and no todos",

			service: &MockService{
				err: fmt.Errorf("Something Bad Happened"),
			},
			inputMethod: "GET",
			inputURL: "http://localhost:8080/todos/",

			expectedStatus: 500,
			// expectedTodos: []entities.Todo{},
		},
		{
			name: "Wrong path gives 404 status and no todos",
			service: &MockService{
				todos: dummyTodos,
			},
			inputMethod:"GET",
			inputURL: "http://localhost:8080/foo",
			
			expectedStatus: 404,
		},
		{
			name: "Wrong path gives 404 status and no todos",
			service: &MockService{
				todos: dummyTodos,
			},
			inputMethod:"GET",
			inputURL: "http://localhost:8080/bar",
			
			expectedStatus: 404,
		}, 
	}

	tests := append(tests, getDisallowedMethodTests()...)
	

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			expect := expectate.Expect(t)

			// service := &MockService{
			// 	err: fmt.Errorf("Something Bad Happened"),	
			// }
		
			w := httptest.NewRecorder()
			r := httptest.NewRequest(test.inputMethod, test.inputURL, nil)
			// r := httptest.NewRequest(http.MethodGet, "http://mywebsite.com", nil)
			
			server := ui.NewHTTP()
		
			server.UseService(test.service)
			server.ServeHTTP(w, r)

			var body []entities.Todo
			json.NewDecoder(w.Result().Body).Decode(&body)
		
			expect(w.Result().StatusCode).ToBe(test.expectedStatus)
			expect(body).ToEqual(test.expectedTodos)
		})
	}
}

func getDisallowedMethodTests() []HTTPTest{
	tests := []HTTPTest{}

	dissallowedMethods := []string {
		http.MethodDelete,
		http.MethodHead,
		http.MethodOptions,
		http.MethodPatch,
		http.MethodPost,
		http.MethodPut,
	}

	for _, method := range dissallowedMethods {
		tests := append(tests, HTTPTest{
			name: fmt.Sprintf("Method %s gives 405 status and no todos", method),

			service:  &MockService{todos: dummyTodos},
			inputURL: "http://localhost:8080/todos/",
			inputMethod: method,

			expectedStatus: http.StatusMethodNotAllowed,
		})
	}
	return tests
}
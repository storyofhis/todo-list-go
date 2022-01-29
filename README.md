# todo-list-go
	
errors on : 
  
```go
tests := append(tests, getDisallowedMethodTests()...)
```
  
```go
for _, method := range dissallowedMethods {
		tests := append(tests, HTTPTest{
			name: fmt.Sprintf("Method %s gives 405 status and no todos", method),

			service:  &MockService{todos: dummyTodos},
			inputURL: "http://localhost:8080/todos/",
			inputMethod: method,

			expectedStatus: http.StatusMethodNotAllowed,
		})
	}
```

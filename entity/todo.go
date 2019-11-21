package entity

// Todo task struct
type Todo struct {
	Title       string
	Description string
	Done        bool
}

// NewTodo empty new task
func NewTodo() Todo {
	return Todo{}
}

// NewTodoWithOptions new task with custom values
func NewTodoWithOptions(title string, description string, done bool) Todo {
	return Todo{title, description, done}
}

// NewTodoDefault new task with default values
func NewTodoDefault() *Todo {
	todo := NewTodoWithOptions(
		"Título",
		"Descrição",
		false,
	)
	return &todo
}

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

// NewTodoNotDone new task with title and description only
func NewTodoNotDone(title string, description string) Todo {
	return Todo{title, description, false}
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

// ToggleDone switch the done attr
func (todo *Todo) ToggleDone() {
	todo.Done = !todo.Done
}

package entity

// TodoList global todolist
var TodoList []Todo

// Todo task struct
type Todo struct {
	ID          int    `json:"id,omitempty" bson:"id,omitempty"`
	Title       string `json:"title,omitempty" bson:"title,omitempty"`
	Description string `json:"description,omitempty" bson:"description,omitempty"`
	Done        bool   `json:"done" bson:"done"`
}

// NewTodo empty new task
func NewTodo() Todo {
	return Todo{}
}

// NewTodoNotDone new task with title and description only
func NewTodoNotDone(title string, description string) Todo {
	return Todo{-1, title, description, false}
}

// NewTodoWithOptions new task with custom values
func NewTodoWithOptions(title string, description string, done bool) Todo {
	return Todo{-1, title, description, done}
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

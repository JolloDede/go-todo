package src

type Todo struct {
	Id    int16
	Title string
	Done  bool
}

func AddTodo(title string) {
	AddTodoToFile(title)
}

func GetTodos() []Todo {
	todos, err := loadTodosFromFile()

	if err != nil {
		return nil
	}

	return todos
}

func CompleteTodo(id int16) {
	updateTodoInFile(int16(id))
}

func DeleteTodo(id int16) {
	deleteTodoInFile(id)
}

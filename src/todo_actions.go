package src

type Todo struct {
	Id    int16
	Title string
}

func AddTodo(title string) {
	writeTodosToFile([]string{title})
}

func GetTodos() []Todo {
	todos, err := loadTodosFromFile()

	if err != nil {
		return nil
	}

	return todos
}

package src

type Todo struct {
	id    int16
	title string
}

func AddTodo(title string) {
	// todos, err := loadTodosFromFile()

	// if err != nil {
	// 	fmt.Println("Cant get Todos from file")
	// 	return
	// }

	// fmt.Println(todos)

	writeTodosToFile([]string{title})
}

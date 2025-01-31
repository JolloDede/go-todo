package src

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Todo struct {
	id    int16
	title string
}

func AddTodo(title string) {
	f, err := os.OpenFile("test.csv", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)

	if err != nil {
		return
	}
	defer f.Close()

	reader := csv.NewReader(f)
	var data []Todo

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error reading CSV data:", err)
			break
		}
		fmt.Println(record)
		todoId, err := strconv.Atoi(record[0])
		if err != nil {
			break
		}
		data = append(data, Todo{id: int16(todoId), title: record[1]})
	}

	w := csv.NewWriter(f)

	w.Write([]string{"1", title})

	w.Flush()
}

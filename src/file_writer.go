package src

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

func loadTodosFromFile() ([]Todo, error) {
	f, err := os.OpenFile("test.csv", os.O_RDONLY, 0644)

	if err != nil {
		return nil, err
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

		todoId, err := strconv.Atoi(record[0])
		if err != nil {
			break
		}
		data = append(data, Todo{Id: int16(todoId), Title: record[1]})
	}
	return data, nil
}

func getLastId() int16 {
	todos, err := loadTodosFromFile()

	if err != nil {
		return 1
	}

	if len(todos) == 0 {
		return 1
	} else {
		return todos[len(todos)-1].Id + 1
	}
}

func writeTodosToFile(todos []string) error {
	f, err := os.OpenFile("test.csv", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)

	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)

	newId := getLastId()

	for i := 0; i < len(todos); i++ {
		w.Write([]string{strconv.Itoa(int(newId)), todos[i]})
		newId++
	}

	w.Flush()
	return nil
}

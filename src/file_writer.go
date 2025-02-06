package src

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
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

		todoCompletion, err := strconv.ParseBool(record[2])
		if err != nil {
			break
		}

		data = append(data, Todo{Id: int16(todoId), Title: record[1], Done: todoCompletion})
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

func AddTodoToFile(todo string) error {
	f, err := os.OpenFile("test.csv", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)

	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)

	newId := getLastId()

	writeTodosToFile(w, [][]string{[]string{strconv.Itoa(int(newId)), todo, strconv.FormatBool(false)}})

	w.Flush()
	return nil
}

func writeTodosToFile(w *csv.Writer, tl [][]string) {
	for i := 0; i < len(tl); i++ {
		w.Write([]string{tl[i][0], tl[i][1], tl[i][2]})
	}
}

func updateTodoInFile(id int16) error {
	f, err := os.OpenFile("test.csv", os.O_RDWR, 6440)

	if err != nil {
		return err
	}
	defer f.Close()

	var nfc [][]string

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")

		if line[0] == strconv.FormatInt(int64(id), 2) {
			state := line[len(line)-1]

			if state == "true" {
				line[len(line)-1] = "false"
			} else {
				line[len(line)-1] = "true"
			}
		}

		nfc = append(nfc, line)
	}

	// TODO handle the errors
	// clear the file
	f.Truncate(0)
	f.Seek(0, 0)

	w := csv.NewWriter(f)

	writeTodosToFile(w, nfc)

	w.Flush()

	return nil
}

func deleteTodoInFile(id int16) error {
	f, err := os.OpenFile("test.csv", os.O_RDWR, 6440)

	if err != nil {
		return err
	}
	defer f.Close()

	var nfc [][]string

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")

		if line[0] != strconv.FormatInt(int64(id), 2) {
			nfc = append(nfc, line)
		}
	}

	// TODO handle the errors
	// clear the file
	f.Truncate(0)
	f.Seek(0, 0)

	w := csv.NewWriter(f)

	writeTodosToFile(w, nfc)

	w.Flush()

	return nil
}

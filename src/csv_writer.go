package src

import (
	"encoding/csv"
	"os"
)

func AddTodo(title string) {
	w := csv.NewWriter(os.Stdout)

	w.Write([]string{title})

	w.Flush()
}

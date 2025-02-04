/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"

	"github.com/JolloDede/todo_go/src"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Returns a tabbed list of your Todos",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		todos := src.GetTodos()
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 4, ' ', 0)

		fmt.Fprintln(w, "ID\tTitle\tDone")
		for i := 0; i < len(todos); i++ {
			fmt.Fprintln(w, strconv.Itoa(int(todos[i].Id))+"\t"+todos[i].Title+"\t"+strconv.FormatBool(todos[i].Done))
		}

		w.Flush()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

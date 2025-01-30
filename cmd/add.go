/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/JolloDede/todo_go/src"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [title]",
	Short: "Add a todo to the list with the [title]",
	Long: `Add a todo to the list of todos.
	The title is the new identifier and a Id thats incremental gets added.`,

	Run: func(cmd *cobra.Command, args []string) {
		// TODO: write to storage
		title := args[0]
		src.AddTodo(title)
		fmt.Println("A new Todo has been created.")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// addCmd.PersistentFlags().StringVarP(&title, "title", "Todo title", "Title of the newly created todo")
	// viper.BindPFlag("title", addCmd.PersistentFlags().Lookup("title"))
}

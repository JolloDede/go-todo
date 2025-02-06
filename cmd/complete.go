/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/JolloDede/todo_go/src"
	"github.com/spf13/cobra"
)

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:   "complete <taskid>",
	Short: "A marks a Todo as done.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		taskid, err := strconv.ParseInt(args[0], 0, 16)
		if err != nil {
			fmt.Println("There was a error converting the id to a number")
			return
		}
		src.CompleteTodo(int16(taskid))
		fmt.Println("complete called")
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// completeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// completeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

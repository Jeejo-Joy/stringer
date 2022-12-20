/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// addCmd represents the add command
var updateCmd2 = &cobra.Command{
	Use:   "update2",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			// todoName := args[0]
			// description, _ := cmd.Flags().GetString("description")
			// deadline, _ := cmd.Flags().GetString("deadline")
			// todo := Todo{
			// 	Name:        todoName,
			// 	Description: description,
			// 	Deadline:    deadline,
			// }

			// todos.Todos = append(todos.Todos, todo)
			// viper.Set("to", todos.Todos)
			// err := viper.WriteConfig()
			// if err != nil {
			// 	fmt.Println(err)
			// 	return
			// }
		} else {
			fmt.Println("Enter just one todo ")
		}
	},
}

func init() {
	fmt.Println(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	fmt.Println("update  init")
	today := time.Now()
	tomorrow := today.AddDate(0, 0, 1)
	updateCmd2.Flags().String("description", "", "Todo description")
	updateCmd2.Flags().String("deadline", tomorrow.String(), "Deadline for the todo")

	rootCmd.AddCommand(updateCmd2)
	fmt.Println("update  init222")
	// today := time.Now()
	// tomorrow := today.AddDate(0, 0, 1)
	// updateCmd2.Flags().String("description", "", "Todo description")
	// updateCmd2.Flags().String("deadline", tomorrow.String(), "Deadline for the todo")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

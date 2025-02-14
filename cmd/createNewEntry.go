/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// createNewEntryCmd represents the createNewEntry command
var createNewEntryCmd = &cobra.Command{
	Use:   "createNewEntry",
	Short: "Create a new time entry in Harvest",
	Long:  `Create a new time entry in Harvest`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("createNewEntry called")
	},
}

func init() {
	rootCmd.AddCommand(createNewEntryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createNewEntryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createNewEntryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

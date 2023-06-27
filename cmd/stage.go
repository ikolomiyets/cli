/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// stageCmd represents the command command
var stageCmd = &cobra.Command{
	Use:   "stage",
	Short: "Stage operations",
	Long:  `A command to run`,
}

var stageListCmd = &cobra.Command{
	Use:   "ls",
	Short: "List stages",
	Long:  "List registered stages",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("listing stages")
	},
}

func init() {
	rootCmd.AddCommand(stageCmd)
	stageCmd.AddCommand(stageListCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// stageCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// stageCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

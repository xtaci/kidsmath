/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "generate addition quiz",
	Run: func(cmd *cobra.Command, args []string) {
		m, n := _parsePattern(cmd.Flags())
		maxdev := _parseMaxdev(cmd.Flags())
		results := generatePrimitive("+", 100, n, m, maxdev)
		for k := range results {
			fmt.Printf("%v = \n", results[k])
		}
	},
}

func init() {
	basicCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

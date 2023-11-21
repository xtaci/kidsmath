/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// basicCmd represents the basic command
var basicCmd = &cobra.Command{
	Use:   "basic",
	Short: "basic math",
	Long:  `generate random (+ - * /) math with number`,
	Run: func(cmd *cobra.Command, args []string) {
		m, n := _parsePattern(cmd.Flags())
		maxdev := _parseMaxdev(cmd.Flags())
		results := generatePrimitive("", 100, n, m, maxdev)
		for k := range results {
			fmt.Printf("%v = \n", results[k])
		}
	},
}

func init() {
	rootCmd.AddCommand(basicCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	//basicCmd.PersistentFlags().String("pattern", "1x2", "defines the pattern of binary-operations")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// basicCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

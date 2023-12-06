/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// subCmd represents the sub command
var subCmd = &cobra.Command{
	Use:   "sub",
	Short: "generate substraction quiz",
	Run: func(cmd *cobra.Command, args []string) {
		m, n := _parsePattern(cmd.Flags())
		maxdev := _parseMaxdev(cmd.Flags())
		results := generatePrimitive("-", 100, n, m, maxdev)
		for k := range results {
			fmt.Printf("%v = \n", results[k])
		}
	},
}

func init() {
	basicCmd.AddCommand(subCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// subCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// subCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

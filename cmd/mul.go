/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// mulCmd represents the mul command
var mulCmd = &cobra.Command{
	Use:   "mul",
	Short: "generate multiplication quiz",
	Run: func(cmd *cobra.Command, args []string) {
		m, n := _parsePattern(cmd.Flags())
		results := generatePrimitive("*", 100, n, m)
		for k := range results {
			fmt.Printf("%v = \n", results[k])
		}

	},
}

func init() {
	basicCmd.AddCommand(mulCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mulCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mulCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

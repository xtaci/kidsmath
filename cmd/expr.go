/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// exprCmd represents the poly command
var exprCmd = &cobra.Command{
	Use:   "expr",
	Short: "generate a random expression , eg: (1+2) * 3",
	Run: func(cmd *cobra.Command, args []string) {
		level, err := cmd.Flags().GetInt("level")
		if err != nil {
			panic(err)
		}
		m, n := _parsePattern(cmd.Flags())
		maxdev := _parseMaxdev(cmd.Flags())
		quizs := generatePrimitive("", 100, m, n, maxdev)
		generateExpr(quizs, level, m, n, maxdev)
		for k := range quizs {
			fmt.Printf("%v = \n", quizs[k])
		}

	},
}

func init() {
	rootCmd.AddCommand(exprCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// polyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// polyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	exprCmd.PersistentFlags().Int("level", 1, "nestedlevel")
}

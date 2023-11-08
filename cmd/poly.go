/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// polyCmd represents the poly command
var polyCmd = &cobra.Command{
	Use:   "poly",
	Short: "generate a random polynomial",
	Run: func(cmd *cobra.Command, args []string) {
		level, err := cmd.Flags().GetInt("level")
		if err != nil {
			panic(err)
		}
		quizs := generate("", 100)
		polyGenerate(quizs, level)
		for k := range quizs {
			fmt.Printf("%v = \n", quizs[k])
		}

	},
}

func init() {
	rootCmd.AddCommand(polyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// polyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// polyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	polyCmd.PersistentFlags().Int("level", 1, "nestedlevel")
}

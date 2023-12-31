/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os"
	"regexp"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kidsmath",
	Short: "A simple program to generate math quizs",
	Long:  `Generate multiplcation, substraction, addition, division quizs within a given range.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.priv.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().String("pattern", "100x10", "defines the boundary of binary-operations")
	rootCmd.PersistentFlags().Int8("maxdev", 50, "defines max deviation in percentile of the generated number, 1-99")
}

func _parsePattern(flags *pflag.FlagSet) (m int, n int) {
	pattern, _ := flags.GetString("pattern")
	patternMatcher := regexp.MustCompile(`([0-9]*)x([0-9]*)`)
	matches := patternMatcher.FindStringSubmatch(pattern)

	n, err := strconv.Atoi(matches[1])
	if err != nil {
		panic(err)
	}

	m, err = strconv.Atoi(matches[2])
	if err != nil {
		panic(err)
	}

	log.Printf("N=%v, M=%v", n, m)
	return m, n
}

func _parseMaxdev(flags *pflag.FlagSet) (maxdev int8) {
	maxdev, _ = flags.GetInt8("maxdev")
	if maxdev > 99 {
		maxdev = 99
	}

	if maxdev < 1 {
		maxdev = 1
	}

	return maxdev
}

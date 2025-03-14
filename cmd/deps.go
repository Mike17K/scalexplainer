/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	golinker "github.com/Mike17K/scalexplainer/src/linker/go"
	pylinker "github.com/Mike17K/scalexplainer/src/linker/python"
	tslinker "github.com/Mike17K/scalexplainer/src/linker/typescript"
	"github.com/spf13/cobra"
)

// depsCmd represents the deps command
var depsCmd = &cobra.Command{
	Use:   "deps",
	Short: "Generates files dependencies json for your codebace",
	Long: `Dependent the language you use for the codebase we run the appropriate command to generate the dependencies file.

	Requirments:
	args:
		- p: path to the codebase
		- l: language of the codebase, one of the following: py, go, ts, js

	Supported languages:
		python: pydeps
		golang: golist
		typescript: madge

	More languages will be added soon.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("deps called")

		// check if the language is supported
		language, _ := cmd.Flags().GetString("language")
		if language != "py" && language != "go" && language != "ts" && language != "js" {
			fmt.Println("Language not supported")
			return
		}

		// get the path
		path, _ := cmd.Flags().GetString("path")
		if path == "" {
			fmt.Println("Path is required")
			return
		}

		// run the appropriate command
		switch language {
		case "py":
			err := pylinker.RunPydeps(path)
			if err != nil {
				fmt.Println("Error running pydeps", err)
				return
			}
		case "go":
			err := golinker.RunGoList(path)
			if err != nil {
				fmt.Println("Error running golist", err)
			}
		case "ts":
			err := tslinker.RunMadge(path)
			if err != nil {
				fmt.Println("Error running madge", err)
			}
		case "js":
			err := tslinker.RunMadge(path)
			if err != nil {
				fmt.Println("Error running madge", err)
			}
		default:
			fmt.Println("Language not supported")
		}
	},
}

func init() {
	rootCmd.AddCommand(depsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// depsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// depsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// add flags
	depsCmd.Flags().StringP("path", "p", "", "path to the codebase")
	depsCmd.Flags().StringP("language", "l", "", "language of the codebase")

	// mark the flags as required
	depsCmd.MarkFlagRequired("path")
	depsCmd.MarkFlagRequired("language")
}

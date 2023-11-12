package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/zcubbs/mq-watch/cmd/cli/cmd/mock"
	"os"
)

var (
	Version string
	Commit  string
	Date    string
)

var (
	rootCmd = &cobra.Command{
		Use:   "",
		Short: "",
		Long:  "",
	}

	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(getVersion())
		},
	}

	aboutCmd = &cobra.Command{
		Use:   "about",
		Short: "Print information about the CLI",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			About()
		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose output")
	rootCmd.DisableSuggestions = true

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(aboutCmd)
	rootCmd.AddCommand(mock.Cmd)
}

func About() {
	fmt.Println("mq-watch is a CLI utility that manages mq-watch server")
	fmt.Println(getFullVersion())
	fmt.Println(getDescription())
	fmt.Println("Author: zakaria.elbouwab")
	fmt.Println("License: MIT")
	fmt.Println("Repository: https://github.com/zcubbs/mq-watch")
}

func getVersion() string {
	return fmt.Sprintf("v%s", Version)
}

func getFullVersion() string {
	return fmt.Sprintf(`
Version: v%s
Commit: %s
Date: %s
`, Version, Commit, Date)
}

func getDescription() string {
	return `
mq-watch is a CLI utility that manages mq-watch server.
`
}

/*
Copyright © 2023 Tobias Grotheer <tobias@grotheer-web.de>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const (
	version = "0.0.2"
)

var showVersion bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "rest-api-cli",
	Short: "A CLI for accessing a rest api.",
	Long: `With this CLI you can access a rest api. Its main purpose is the check command, to run
Nagio compatible checks agains a rest api endpoint.`,
	Run: runRoot,
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
	rootCmd.Flags().BoolVarP(&showVersion, "version", "V", false, "Show the version")
}

func runRoot(cmd *cobra.Command, args []string) {
	if showVersion {
		fmt.Printf("rest-api-cli version %s\n", version)
	}
	os.Exit(0)
}

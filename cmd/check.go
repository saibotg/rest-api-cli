/*
Copyright © 2023 Tobias Grotheer <tobias@grotheer-web.de>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Run a check agains the rest api",
	Long:  `Run a nagios compatible check against the api endpoint.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("check called")
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
}

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of rcli",
	Long:  `Version of the rcli - multipurpose CLI`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("rcli - multipurpose CLI")
	},
}

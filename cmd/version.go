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
	Short: "Print the version number of harborclient",
	Long:  `All software has versions. This is harborclient`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("App HarborClient Version: v1.0 -- HEAD")
	},
}

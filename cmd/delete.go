package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"harborclient/cmd/delete"
	"harborclient/logger"
)

var cmdDelete = &cobra.Command{
	Use:   "delete [command]",
	Short: "use Delete method to Delete harbor repo tag",
	Long: `
    `,
	Args: cobra.MinimumNArgs(1),
	Run:  deleteRun,
}

func deleteRun(cmd *cobra.Command, args []string) {
	checkDeleteArgs()
}

// Delete flags
var rName string

func checkDeleteArgs() {
	logger.Logging.WithFields(logrus.Fields{"resourceName": resourceName}).Info("Delete input args:")
}

func init() {
	// local Flags
	cmdDelete.Flags().StringVar(&rName, "resourceName", "", "Delete all projects")

	//required Flags
	cmdDelete.MarkFlagRequired("resourceName")

	//reg command
	rootCmd.AddCommand(cmdDelete)
	//reg sub command
	cmdDelete.AddCommand(delete.CmdTag)
}

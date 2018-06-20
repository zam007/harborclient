package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"harborclient/cmd/get"
	"harborclient/logger"
)

var cmdGet = &cobra.Command{
	Use:   "get [command]",
	Short: "use get method to get harbor repo info",
	Long: `
    `,
	Args: cobra.MinimumNArgs(1),
	Run:  getRun,
}

func getRun(cmd *cobra.Command, args []string) {
	checkGetArgs()
}

// get flags
var resourceName string

func checkGetArgs() {
	logger.Logging.WithFields(logrus.Fields{"resourceName": resourceName}).Info("get input args:")
}

func init() {
	// local Flags
	cmdGet.Flags().StringVar(&resourceName, "resourceName", "", "get all projects")

	//required Flags
	//cmdGet.MarkFlagRequired("resourceName")

	//reg command
	rootCmd.AddCommand(cmdGet)
	//reg sub command
	cmdGet.AddCommand(get.CmdProjects)
	cmdGet.AddCommand(get.CmdRepo)
	cmdGet.AddCommand(get.CmdTag)
}

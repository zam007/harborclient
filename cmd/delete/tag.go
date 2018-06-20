package delete

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"harborclient/cmd/sdk"
	"harborclient/logger"
)

var CmdTag = &cobra.Command{
	Use:   "tag --flags",
	Short: "return tag name by assign harbor project",
	Run:   tagRun,
}

var keepTagNum int

func tagRun(cmd *cobra.Command, args []string) {
	checkArgs()

	repos := sdk.HClient.GetRepoName(sdk.HttpClient)

	for _, r := range repos {
		tags := sdk.HClient.GetTag(sdk.HttpClient, r)
		if len(tags) == 0 {
			logger.Logging.Warn("get ", r, " tags failed")
			continue
		}

		if len(tags) > keepTagNum && keepTagNum >= 30 {
			logger.Logging.Info("repo : ", r)
			logger.Logging.Info("tags : ", tags)
			logger.Logging.Info("will delete :", tags[:len(tags)-keepTagNum])

			for _, t := range tags[:len(tags)-keepTagNum] {
				sdk.HClient.DeleteRepoTag(sdk.HttpClient, r, t)
			}
		} else {
			logger.Logging.WithFields(logrus.Fields{"repo": r, "tagCount": len(tags)}).Info("Do nothing")
		}
	}
}

var repoName string

func init() {
	// get harbor http client
	sdk.HttpClient = sdk.HClient.Login()

	CmdTag.Flags().StringVar(&repoName, "repoName", "", "get tag from assign harbor project")
	CmdTag.Flags().IntVarP(&keepTagNum, "keepTagNum", "n", 0, "designated the tag keep number ")

}

func checkArgs() {
	if keepTagNum <= 30 {
		logger.Logging.Info("no number input ,use default : ", 30)
		keepTagNum = 30
	} else {
		logger.Logging.WithFields(logrus.Fields{"keepTagNum": keepTagNum}).Info("get number from input")
	}
}

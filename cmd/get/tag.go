package get

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"harborclient/cmd/sdk"
	"os"
)

var CmdTag = &cobra.Command{
	Use:   "tag ",
	Short: "",
	Run:   getTagRun,
}

type backData struct {
	Text       string
	IsMetadata bool
}

func getTagRun(cmd *cobra.Command, args []string) {
	checkGetTagRunArgs()

	tags := sdk.HClient.GetTag(sdk.HttpClient, repoName)

	var backDatas []backData
	if len(tags) >= 5 {
		for _, t := range tags[len(tags)-5:] {

			var backData backData
			backData.Text = t
			backData.IsMetadata = false

			backDatas = append(backDatas, backData)
		}
	} else {
		os.Exit(1)
	}

	jsonMap := make(map[string][]backData)
	jsonMap["data"] = backDatas

	jsonString, _ := json.Marshal(jsonMap)
	fmt.Println(string(jsonString))

}

var repoName string

func init() {
	// get harbor http client
	sdk.HttpClient = sdk.HClient.Login()

	CmdTag.Flags().StringVar(&repoName, "repoName", "", "get tag from assign harbor project")

}

func checkGetTagRunArgs() {

}

package get

import (
	"fmt"
	"github.com/spf13/cobra"
	"harborclient/cmd/sdk"
)

var CmdRepo = &cobra.Command{
	Use:   "repo",
	Short: "return repo name by assign harbor project",
	Run:   RepoRun,
}

func RepoRun(cmd *cobra.Command, args []string) {
	targetUrl := fmt.Sprintf("%s://%s/api/repositories?project_id=%d", sdk.HClient.Protocol, sdk.HClient.Host, projectId)
	client := sdk.HClient.Login()

	resp, err := client.Get(targetUrl)
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}

	rep := sdk.HClient.GetRepoName(sdk.HttpClient)
	for _, r := range rep {
		fmt.Println(r)
	}

}

var projectId int

func init() {
	// get harbor http client
	sdk.HttpClient = sdk.HClient.Login()

	// local Flags
	CmdRepo.Flags().IntVar(&projectId, "projectid", 0, "get repo from assign harbor project")

	//required Flags
	//CmdRepo.MarkFlagRequired("projectid")

}

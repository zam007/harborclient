package get

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"harborclient/cmd/sdk"
	"io/ioutil"
)

var CmdProjects = &cobra.Command{
	Use:   "projects",
	Short: "return all harbor project",
	Run:   ProjectsRun,
}

func ProjectsRun(cmd *cobra.Command, args []string) {
	targetUrl := fmt.Sprintf("%s://%s/api/projects", sdk.HClient.Protocol, sdk.HClient.Host)
	resp, err := sdk.HttpClient.Get(targetUrl)
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}

	if data, err := ioutil.ReadAll(resp.Body); err == nil {
		json.Unmarshal(data, &sdk.Projects)
		for _, p := range sdk.Projects {
			fmt.Println(p.Name)
		}
	}
}

func init() {
	// get harbor http client
	sdk.HttpClient = sdk.HClient.Login()
}

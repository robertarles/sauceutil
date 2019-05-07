// Copyright © 2019 Robert Arles robert@arles.us
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

// getJobsCmd represents the getJobs command
var getjobsCmd = &cobra.Command{
	Use:   "getjobs 5",
	Short: "Retrieve a list of the most recent jobs run.",
	Long:  `TODO: Long version -> Retrieve a list of the most recent jobs run.`,
	Run: func(cmd *cobra.Command, args []string) {
		var maxJobs string
		if len(args) == 1 {
			maxJobs = args[0]
		} else {
			fmt.Printf("upload requires a parameter to specify how many jobs to list\ntry the --help option\n")
			os.Exit(1)
		}
		GetJobs(maxJobs)
	},
}

func init() {
	rootCmd.AddCommand(getjobsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getJobsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getJobsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type customData struct {
	BuildNumber      string `json:"BUILD_NUMBER"`
	JenkinsBuildName string `json:"JENKINS_BUILD_NAME"`
	GitCommit        string `json:"GIT_COMMIT"`
}

type jobData struct {
	BrowserShortVersion string `json:"browser_short_version"`
	VideoURL            string `json:"video_url"`
	CreationTime        int64  `json:"creation_time"`
	//CustomData            customData `json:"custom-data"` // TODO: handle the saucelabs  error? sometimes returns CustomData.BuildNumber as an int, not string
	BrowserVersion        string   `json:"browser_version"`
	Owner                 string   `json:"owner"`
	ID                    string   `json:"id"`
	Container             bool     `json:"container"`
	RecordScreenshots     bool     `json:"record_screenshots"`
	RecordVideo           bool     `json:"record_video"`
	Build                 string   `json:"build"`
	Passed                bool     `json:"passed"`
	Public                string   `json:"public"`
	EndTime               int64    `json:"end_time"`
	Status                string   `json:"status"`
	LogURL                string   `json:"log_url"`
	StartTime             int64    `json:"start_time"`
	Proxied               bool     `json:"proxied"`
	ModificationTime      int64    `json:"modification_time"`
	Tags                  []string `json:"tags"`
	Name                  string   `json:"name"`
	CommandsNotSuccessful uint32   `json:"commands_not_successful"`
	ConsolidatedStatus    string   `json:"consolidated_stats"`
	AssignedTunnelID      string   `json:"assigned_tunnel_id"`
	Error                 string   `json:"error"`
	OS                    string   `json:"os"`
	Breakpointed          bool     `json:"breakpointed"`
	Browser               string   `json:"browser"`
}

// GetJobs Get details for [count] last jobs
func GetJobs(count string) ([]jobData, error) {
	var retVal []jobData
	username := os.Getenv("SAUCE_USERNAME")
	accessKey := os.Getenv("SAUCE_ACCESS_KEY")

	client := &http.Client{}
	request, err := http.NewRequest("GET", apiURL+"/jobs?limit="+count+"&full=true", nil)
	request.SetBasicAuth(username, accessKey)
	response, err1 := client.Do(request)
	if err1 != nil {
		fmt.Printf("The http request failed with error %s\n", err)
		return retVal, err1
	}
	// success path
	retVal = []jobData{}
	data, _ := ioutil.ReadAll(response.Body)
	err2 := json.Unmarshal(data, &retVal)
	fmt.Println(string(data))
	return retVal, err2
}

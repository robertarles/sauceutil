// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
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

// getJobCmd represents the getJob command
var getJobCmd = &cobra.Command{
	Use:   "getJob",
	Short: "Get details on a specific job",
	Long:  `TODO: long description -> Get details on a specific job`,
	Run: func(cmd *cobra.Command, args []string) {
		var jobID string
		if len(args) == 1 {
			jobID = args[0]
		} else {
			fmt.Printf("upload requires a jobID parameter (e.g. from getJobs command)\ntry the --help option\n")
			os.Exit(1)
		}
		GetJob(jobID)
	},
}

func init() {
	rootCmd.AddCommand(getJobCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getJobCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getJobCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// type jobData struct {
// 	BrowserShortVersion string `json:"browser_short_version"`
// 	VideoURL            string `json:"video_url"`
// 	CreationTime        int64  `json:"creation_time"`
// 	//CustomData            customData `json:"custom-data"` // TODO: handle the saucelabs  error? sometimes returns CustomData.BuildNumber as an int, not string
// 	BrowserVersion        string   `json:"browser_version"`
// 	Owner                 string   `json:"owner"`
// 	ID                    string   `json:"id"`
// 	Container             bool     `json:"container"`
// 	RecordScreenshots     bool     `json:"record_screenshots"`
// 	RecordVideo           bool     `json:"record_video"`
// 	Build                 string   `json:"build"`
// 	Passed                bool     `json:"passed"`
// 	Public                string   `json:"public"`
// 	EndTime               int64    `json:"end_time"`
// 	Status                string   `json:"status"`
// 	LogURL                string   `json:"log_url"`
// 	StartTime             int64    `json:"start_time"`
// 	Proxied               bool     `json:"proxied"`
// 	ModificationTime      int64    `json:"modification_time"`
// 	Tags                  []string `json:"tags"`
// 	Name                  string   `json:"name"`
// 	CommandsNotSuccessful uint32   `json:"commands_not_successful"`
// 	ConsolidatedStatus    string   `json:"consolidated_stats"`
// 	AssignedTunnelID      string   `json:"assigned_tunnel_id"`
// 	Error                 string   `json:"error"`
// 	OS                    string   `json:"os"`
// 	Breakpointed          bool     `json:"breakpointed"`
// 	Browser               string   `json:"browser"`
// }

// GetJob Get detail on the specific job ID
func GetJob(jobID string) {

	username := os.Getenv("SAUCE_USERNAME")
	accessKey := os.Getenv("SAUCE_ACCESS_KEY")

	client := &http.Client{}
	request, err := http.NewRequest("GET", apiURL+"/jobs/"+jobID, nil)
	request.SetBasicAuth(username, accessKey)
	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("The http request failed with error %s\n", err)
	} else {
		respBody := jobData{}
		data, _ := ioutil.ReadAll(response.Body)
		json.Unmarshal(data, &respBody)
		fmt.Println(string(data))
		//fmt.Printf("git commit -> %s", respBody.CustomData.GitCommit)
	}

}

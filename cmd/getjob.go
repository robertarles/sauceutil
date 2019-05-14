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

var getJobID string

// getJobCmd represents the getJob command
var getjobCmd = &cobra.Command{
	Use:   "getjob -j {jobID}",
	Short: "Get details on a specific job",
	Long:  `Get details on a specific job`,
	Run: func(cmd *cobra.Command, args []string) {
		var _, jsonString, err = GetJob(getJobID)
		if err == nil {
			fmt.Printf("%s\n", jsonString)
		} else {
			fmt.Printf("%s\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(getjobCmd)

	// Here you will define your flags and configuration settings.
	getjobCmd.Flags().StringVarP(&getJobID, "jobid", "j", "", "Saucelabs Job ID")
	getjobCmd.MarkFlagRequired("jobid")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getJobCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getJobCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// GetJob Get detail on the specific job ID
func GetJob(jobID string) (jobData JobData, jsonString string, err error) {

	username := os.Getenv("SAUCE_USERNAME")
	accessKey := os.Getenv("SAUCE_ACCESS_KEY")

	client := &http.Client{}
	request, err := http.NewRequest("GET", apiURL+"/jobs/"+jobID, nil)
	request.SetBasicAuth(username, accessKey)
	response, err := client.Do(request)
	if err != nil {
		return JobData{}, fmt.Sprintf(`{"error": "The http request failed with error %s}"`, err), err
	}
	respBody := JobData{}
	data, _ := ioutil.ReadAll(response.Body)
	json.Unmarshal(data, &respBody)
	jsonBytes, _ := json.MarshalIndent(respBody, "", "  ")
	return respBody, string(jsonBytes), nil

}

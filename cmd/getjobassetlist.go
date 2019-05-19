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

var jobID string

// getjobassetlistCmd represents the getjobassetlist command
var getjobassetlistCmd = &cobra.Command{
	Use:   "getjobassetlist -j {JobID}",
	Short: "Get a list of files associated to a job.",
	Long:  `Get a list of files associated to a job.`,
	Run: func(cmd *cobra.Command, args []string) {
		var _, jsonString, err = GetJobAssetList(jobID)
		if err != nil {
			fmt.Printf("%s\n", err)
			os.Exit(1)
		}
		fmt.Printf("%s\n", jsonString)

	},
}

func init() {
	rootCmd.AddCommand(getjobassetlistCmd)
	// Here you will define your flags and configuration settings.

	getjobassetlistCmd.Flags().StringVarP(&jobID, "jobid", "j", "", "The Saucelabs job ID to get an asset list for.")
	getjobassetlistCmd.MarkFlagRequired("jobid")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getjobassetlistCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getjobassetlistCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// GetJobAssetList requests the file assets associated with the job
func GetJobAssetList(jobID string) (responseBody AssetListData, jsonString string, err error) {

	username := os.Getenv("SAUCE_USERNAME")
	accessKey := os.Getenv("SAUCE_ACCESS_KEY")

	client := &http.Client{}
	request, err := http.NewRequest("GET", apiURL+"/"+username+"/jobs/"+jobID+"/assets", nil)
	request.SetBasicAuth(username, accessKey)
	response, err := client.Do(request)

	if err != nil {
		return AssetListData{}, "", err
	}

	data, err := ioutil.ReadAll(response.Body)
	responseBody = AssetListData{}
	json.Unmarshal(data, &responseBody)
	jsonBytes, _ := json.MarshalIndent(responseBody, "", "  ")
	return responseBody, string(jsonBytes), err

}

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
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

// getjobassetlistCmd represents the getjobassetlist command
var getjobassetlistCmd = &cobra.Command{
	Use:   "getjobassetlist {jobID}",
	Short: "Get a list of files associated to a job.",
	Long:  `Get a list of files associated to a job.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires jobID argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		var jobID = args[0]
		var _, jsonString, err = GetJobAssetList(jobID)
		if err == nil {
			fmt.Printf(jsonString)
		}
	},
}

func init() {
	rootCmd.AddCommand(getjobassetlistCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getjobassetlistCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getjobassetlistCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// GetJobAssetList requests the file assets associated with the job
func GetJobAssetList(jobID string) (responseBody assetListData, jsonString string, err error) {

	username := os.Getenv("SAUCE_USERNAME")
	accessKey := os.Getenv("SAUCE_ACCESS_KEY")

	client := &http.Client{}
	request, err := http.NewRequest("GET", apiURL+"/"+username+"/jobs/"+jobID+"/assets", nil)
	request.SetBasicAuth(username, accessKey)
	response, err := client.Do(request)
	jsonString = ""
	if err == nil {
		data, err := ioutil.ReadAll(response.Body)
		jsonString = string(data)
		responseBody := assetListData{}
		json.Unmarshal(data, &responseBody)
		return responseBody, jsonString, err
	}
	return assetListData{}, "", err
}

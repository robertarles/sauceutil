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

// getJobCmd represents the getJob command
var getjobCmd = &cobra.Command{
	Use:   "getjob {jobID}",
	Short: "Get details on a specific job",
	Long:  `Get details on a specific job`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires jobID argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		var jobID = args[0]
		GetJob(jobID)
	},
}

func init() {
	rootCmd.AddCommand(getjobCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getJobCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getJobCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

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

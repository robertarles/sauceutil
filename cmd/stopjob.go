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
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var stopJobID string

// stopjobCmd represents the stopjob command
var stopjobCmd = &cobra.Command{
	Use:   "stopjob -j {jobID}",
	Short: "Terminates a running Saucelabs job",
	Long:  `Terminates a running Saucelabs job`,
	Run: func(cmd *cobra.Command, args []string) {
		var jsonString, err = StopJob(stopJobID)
		if err == nil {
			fmt.Printf("%s\n", jsonString)
		} else {
			fmt.Printf("%s\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(stopjobCmd)

	// Here you will define your flags and configuration settings.
	stopjobCmd.Flags().StringVarP(&stopJobID, "jobid", "j", "", "Saucelabs Job ID")
	stopjobCmd.MarkFlagRequired("jobid")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// stopjobCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// stopjobCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// StopJob Get detail on the specific job ID
func StopJob(jobID string) (jsonString string, err error) {

	// TODO: add check for jobs existence and running status, the test if it's actually stopped
	username := os.Getenv("SAUCE_USERNAME")
	accessKey := os.Getenv("SAUCE_ACCESS_KEY")

	client := &http.Client{}
	request, err := http.NewRequest("PUT", apiURL+"/"+username+"/jobs/"+jobID+"/stop", nil)
	request.SetBasicAuth(username, accessKey)
	response, err := client.Do(request)
	if err != nil {
		return `{"status": "http error"}`, err
	}
	jsonString = fmt.Sprintf(`{"status": "%s", "statusCode:": %d}`, response.Status, response.StatusCode)
	return jsonString, nil

}

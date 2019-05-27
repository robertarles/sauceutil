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
	"strings"

	"github.com/spf13/cobra"
)

var stopJobID string

// TODO: fix this to return an object so that the -o flag can format output

// stopjobCmd represents the stopjob command
var stopjobCmd = &cobra.Command{
	Use:   "stopjob -j {jobID}",
	Short: "Terminates a running Saucelabs job",
	Long:  `Terminates a running Saucelabs job`,
	Run: func(cmd *cobra.Command, args []string) {
		var statusCode, err = StopJob(stopJobID)
		if err != nil {
			fmt.Printf("%s\n", err)
			os.Exit(1)
		}
		fmt.Printf("http %d\n", statusCode)
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
func StopJob(jobID string) (statusCode int, err error) {

	// TODO: add check for jobs existence and running status, then after stop command -> test if it's actually stopped

	username := os.Getenv("SAUCE_USERNAME")
	accessKey := os.Getenv("SAUCE_ACCESS_KEY")

	getJobResponse, _, getJobErr := GetJob(jobID)
	if getJobErr != nil {
		return 0, getJobErr
	}
	if (strings.ToLower(getJobResponse.Status) != "running") && (strings.ToLower(getJobResponse.Status) != "in progress") {
		return 0, fmt.Errorf(fmt.Sprintf("job %s does not appear to be running\n", jobID))
	}

	client := &http.Client{}
	request, reqErr := http.NewRequest("PUT", apiURL+"/"+username+"/jobs/"+jobID+"/stop", nil)
	if reqErr != nil {
		return 0, reqErr
	}
	request.SetBasicAuth(username, accessKey)
	response, doErr := client.Do(request)
	if doErr != nil {
		return 0, doErr
	}
	statusCode = response.StatusCode
	if statusCode > 299 {
		return 0, fmt.Errorf(fmt.Sprintf("stop job request returned status code %d\n", statusCode))
	}

	getJobResponseAft, _, getJobErrAft := GetJob(jobID)
	if getJobErr != nil {
		return 0, getJobErrAft
	}
	if strings.ToLower(getJobResponseAft.Status) == "running" || (strings.ToLower(getJobResponse.Status) != "in progress") {
		return 0, fmt.Errorf(fmt.Sprintf("job %s does not appear to have stopped\n", jobID))
	}

	return statusCode, nil

}

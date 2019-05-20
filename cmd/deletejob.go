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
	"net/http"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var deleteJobID string

// deletejobCmd represents the deletejob command
var deletejobCmd = &cobra.Command{
	Use:   "deletejob",
	Short: "Removes the job from the Saucelabs system with all the linked assets",
	Long:  `Removes the job from the Saucelabs system with all the linked assets`,
	Run: func(cmd *cobra.Command, args []string) {
		var _, jsonString, err = DeleteJob(deleteJobID)
		if err == nil {
			fmt.Printf("%s\n", jsonString)
		} else {
			fmt.Printf("%+v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(deletejobCmd)

	// Here you will define your flags and configuration settings.
	deletejobCmd.Flags().StringVarP(&deleteJobID, "jobid", "j", "", "Saucelabs Job ID")
	deletejobCmd.MarkFlagRequired("jobid")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deletejobCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deletejobCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// DeleteJob Get detail on the specific job ID
func DeleteJob(deleteJobID string) (deleteData DeleteJobData, deleteJSONString string, err error) {

	username := os.Getenv("SAUCE_USERNAME")
	accessKey := os.Getenv("SAUCE_ACCESS_KEY")

	jobData, _, getErr := GetJob(deleteJobID)
	if getErr != nil {
		return DeleteJobData{}, "{}", getErr
	}
	if jobData.ID != deleteJobID {
		return DeleteJobData{}, "", fmt.Errorf(fmt.Sprintf("job %s not found", deleteJobID))
	}
	if strings.ToLower(jobData.Status) == "running" {
		return DeleteJobData{}, "", fmt.Errorf(fmt.Sprintf("job %s is running", deleteJobID))
	}

	client := &http.Client{}
	request, reqErr := http.NewRequest("DELETE", apiURL+"/"+username+"/jobs/"+deleteJobID, nil)
	if reqErr != nil {
		return DeleteJobData{}, "", reqErr
	}
	request.SetBasicAuth(username, accessKey)
	response, doErr := client.Do(request)
	if doErr != nil {
		return DeleteJobData{}, "", err
	}
	if response.StatusCode > 299 {
		return DeleteJobData{}, "", fmt.Errorf(fmt.Sprintf("request error with status code %d\n", response.StatusCode))
	}

	respBody := DeleteJobData{}
	decoder := json.NewDecoder(response.Body)
	decodeErr := decoder.Decode(&respBody)
	if decodeErr != nil {
		return DeleteJobData{}, "", decodeErr
	}
	jsonBytes, marshErr := json.MarshalIndent(respBody, "", "  ")
	if marshErr != nil {
		return DeleteJobData{}, "", marshErr
	}
	return respBody, string(jsonBytes), nil

}

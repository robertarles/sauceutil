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
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var jobID string

// getjobassetlistCmd represents the getjobassetlist command
var getjobassetlistCmd = &cobra.Command{
	Use:   "assetlist -i {JobID}",
	Short: "Get a list of files associated to a job.",
	Long:  `Get a list of files associated to a job.`,
	Run: func(cmd *cobra.Command, args []string) {

		var jsonString, err = GetJobAssetList(jobID)
		if err != nil {
			fmt.Printf("%+v\n", err)
			os.Exit(1)
		}
		if len(OutFormat) == 0 {
			fmt.Printf("%s\n\n", jsonString)
		} else {
			printHeader := true
			err := OPrintFormatted(OutFormat, jsonString, printHeader)
			if err != nil {
				fmt.Printf("%+v\n", err)
				os.Exit(1)
			}
		}
		os.Exit(0)
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

	getjobassetlistCmd.Flags().StringVarP(&jobID, "id", "i", "", "The Saucelabs job ID to get an asset list for.")
	getjobassetlistCmd.MarkFlagRequired("id")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getjobassetlistCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getjobassetlistCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// GetJobAssetList requests the file assets associated with the job
func GetJobAssetList(jobID string) (jsonString string, err error) {

	username := os.Getenv("SAUCE_USERNAME")
	accessKey := os.Getenv("SAUCE_ACCESS_KEY")

	client := &http.Client{}
	request, err := http.NewRequest("GET", apiURL+"/"+username+"/jobs/"+jobID+"/assets", nil)
	request.SetBasicAuth(username, accessKey)
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	jsonBytes, readErr := ioutil.ReadAll(response.Body)
	if readErr != nil {
		return "", readErr
	}
	return string(jsonBytes), nil

}

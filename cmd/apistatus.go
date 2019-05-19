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

// apistatusCmd represents the apistatus command
var apistatusCmd = &cobra.Command{
	Use:   "apistatus",
	Short: "Request the current API status.",
	Long:  `Request the current API status.`,
	Run: func(cmd *cobra.Command, args []string) {
		var _, jsonString, err = GetAPIStatus()
		if err != nil {
			fmt.Printf("%s\n", err)
			os.Exit(1)
		}
		fmt.Printf("%s\n", jsonString)
	},
}

func init() {
	rootCmd.AddCommand(apistatusCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// apistatusCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// apistatusCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// GetAPIStatus Get the status of the Saucelabs API
func GetAPIStatus() (apiStatusResponse APIStatusResponseData, jsonString string, err error) {
	response, err := http.Get(apiURL + "/info/status")

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {

		data, err := ioutil.ReadAll(response.Body)

		var response APIStatusResponseData
		json.Unmarshal(data, &response)

		if err != nil {
			return response, "", err
		}
		jsonString, _ := json.MarshalIndent(response, "", "  ")
		return APIStatusResponseData{}, string(jsonString), nil
	}
	return APIStatusResponseData{}, "", err
}

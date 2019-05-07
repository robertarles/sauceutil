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

	"github.com/spf13/cobra"
)

// apistatusCmd represents the apistatus command
var apistatusCmd = &cobra.Command{
	Use:   "apistatus",
	Short: "Request the current API status.",
	Long:  `TODO: Longer desc`,
	Run: func(cmd *cobra.Command, args []string) {
		GetAPIStatus()
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

type statusResponse struct {
	WaitTime           float32 `json:"wait_time"`
	ServiceOperational bool    `json:"service_operational"`
	StatusMessage      string  `json:"status_message"`
}

// GetAPIStatus Get the status of the Saucelabs API
func GetAPIStatus() {
	response, err := http.Get(apiURL + "/info/status")

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {

		data, _ := ioutil.ReadAll(response.Body)

		response := statusResponse{}
		json.Unmarshal(data, &response)

		fmt.Println()
		//fmt.Println(string(data))
		fmt.Println("WaitTime: ", response.WaitTime)
		fmt.Println("ServiceOperational: ", response.ServiceOperational)
		fmt.Println("StatusMessage: ", response.StatusMessage)
	}
}

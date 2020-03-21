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
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var tunnelID string

// getJobCmd represents the getJob command
var tunnelCmd = &cobra.Command{
	Use:   "tunnel -i {tunnelID}",
	Short: "Get details on a specific tunnel",
	Long:  `Get details on a specific tunnel`,
	Run: func(cmd *cobra.Command, args []string) {
		var tunnelData, jsonString, err = Tunnel(tunnelID)
		if err != nil {
			fmt.Printf("%s\n", err)
			os.Exit(1)
		}
		if len(OutFormat) == 0 {
			fmt.Printf("%s\n\n", jsonString)
		} else {
			printHeader := true
			OPrintStruct(OutFormat, tunnelData, printHeader)
		}
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(tunnelCmd)

	// Here you will define your flags and configuration settings.
	tunnelCmd.Flags().StringVarP(&tunnelID, "id", "i", "", "Saucelabs Tunnel ID")
	tunnelCmd.MarkFlagRequired("id")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getJobCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getJobCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// Tunnel Get detail on the specific job ID
func Tunnel(tunnelID string) (respBody TunnelData, jsonString string, err error) {

	username := os.Getenv("SAUCE_USERNAME")
	accessKey := os.Getenv("SAUCE_ACCESS_KEY")

	client := &http.Client{}
	request, err := http.NewRequest("GET", apiURL+"/"+username+"/tunnels/"+tunnelID, nil)
	request.SetBasicAuth(username, accessKey)
	response, err := client.Do(request)
	if err != nil {
		return TunnelData{}, fmt.Sprintf(`{"error": "The http request failed with error %s}"`, err), err
	}

	respBody = TunnelData{}
	decoder := json.NewDecoder(response.Body)
	decodeErr := decoder.Decode(&respBody)
	if decodeErr != nil {
		return TunnelData{}, "", decodeErr
	}
	jsonBytes, marshErr := json.MarshalIndent(respBody, "", "  ")
	if marshErr != nil {
		return TunnelData{}, "", marshErr
	}

	return respBody, string(jsonBytes), nil

}

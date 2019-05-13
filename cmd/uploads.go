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

// uploadsCmd represents the uploads command
var uploadsCmd = &cobra.Command{
	Use:   "uploads",
	Short: "A list of files already uploaded to sauce-storage.",
	Long:  `A list of files already uploaded to sauce-storage.`,
	Run: func(cmd *cobra.Command, args []string) {
		var _, jsonString, err = Uploads()
		if err != nil {
			fmt.Printf("%s", err)
		}
		fmt.Printf("%s", jsonString)
	},
}

var apiURL = "https://saucelabs.com/rest/v1"

func init() {
	rootCmd.AddCommand(uploadsCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// uploadsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// uploadsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// Uploads returns a list of files in the sauce-storage upload area
func Uploads() (storageResponse StorageResponse, jsonString string, err error) {

	username := os.Getenv("SAUCE_USERNAME")
	accessKey := os.Getenv("SAUCE_ACCESS_KEY")

	client := &http.Client{}
	request, err := http.NewRequest("GET", apiURL+"/storage/"+username, nil)
	request.SetBasicAuth(username, accessKey)
	response, err := client.Do(request)
	if err != nil {
		var jsonString = fmt.Sprintf(`{"error": "The http request failed with error %s"}`, err)
		return StorageResponse{}, jsonString, err
	}
	storageResponse = StorageResponse{}
	data, _ := ioutil.ReadAll(response.Body)
	json.Unmarshal(data, &storageResponse)

	if len(storageResponse.Files) > 0 {
		jsonBytes, _ := json.MarshalIndent(storageResponse, "", "  ")
		return storageResponse, string(jsonBytes), nil
	}
	return StorageResponse{}, `{"message": "No files found."}`, nil

}

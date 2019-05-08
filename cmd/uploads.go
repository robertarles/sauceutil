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
		Uploads()
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

func Uploads() {

	username := os.Getenv("SAUCE_USERNAME")
	accessKey := os.Getenv("SAUCE_ACCESS_KEY")

	client := &http.Client{}
	request, err := http.NewRequest("GET", apiURL+"/storage/"+username, nil)
	request.SetBasicAuth(username, accessKey)
	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("The http request failed with error %s\n", err)
	} else {
		respBody := storageResponse{}
		data, _ := ioutil.ReadAll(response.Body)
		json.Unmarshal(data, &respBody)

		if len(respBody.Files) > 0 {
			fmt.Printf(string(data))
		} else {
			fmt.Println("No files found.")
		}
	}

}

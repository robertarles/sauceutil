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

// uploadsCmd represents the uploads command
var uploadsCmd = &cobra.Command{
	Use:   "uploads",
	Short: "A list of files already uploaded to sauce-storage.",
	Long:  `A list of files already uploaded to sauce-storage.`,
	Run: func(cmd *cobra.Command, args []string) {
		var jsonString, err = Uploads()
		if err != nil {
			fmt.Printf("%s\n", err)
			os.Exit(1)
		}
		if err != nil {
			fmt.Printf("%+v\n", err)
			os.Exit(1)
		}
		if len(OutFormat) == 0 {
			fmt.Printf("%s\n", jsonString)
		} else {
			printHeader := true
			err := OPrintFormatted(OutFormat, jsonString, printHeader)
			if err != nil {
				fmt.Printf("%+v\n", err)
				os.Exit(1)
			}
		}
		os.Exit(0)

		// if err != nil {
		// 	fmt.Printf("%s\n", err)
		// 	os.Exit(1)
		// }
		// fmt.Printf("%s\n", jsonString)
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
func Uploads() (jsonString string, err error) {

	username := os.Getenv("SAUCE_USERNAME")
	accessKey := os.Getenv("SAUCE_ACCESS_KEY")

	client := &http.Client{}
	request, reqErr := http.NewRequest("GET", apiURL+"/storage/"+username, nil)
	if reqErr != nil {
		return "", reqErr
	}
	request.SetBasicAuth(username, accessKey)
	response, doErr := client.Do(request)
	if doErr != nil {
		return "", doErr
	}
	jsonBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(jsonBytes), nil

}

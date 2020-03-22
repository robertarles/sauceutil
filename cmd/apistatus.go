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

// apistatusCmd represents the apistatus command
var apistatusCmd = &cobra.Command{
	Use:   "apistatus",
	Short: "Request the current API status.",
	Long:  `Request the current API status.`,
	Run: func(cmd *cobra.Command, args []string) {
		var jsonString, err = GetAPIStatus()
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
func GetAPIStatus() (jsonString string, err error) {

	response, err := http.Get(apiURL + "/info/status")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return "", err
	}
	defer response.Body.Close()
	body, errReading := ioutil.ReadAll(response.Body)
	if errReading != nil {
		return "", fmt.Errorf("error reading API response")
	}
	return string(body), nil

}

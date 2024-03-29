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

var maxJobs uint

// getJobsCmd represents the getJobs command
var getjobsCmd = &cobra.Command{
	Use:   "jobs -m {maxJobs} (defaults to 5)",
	Short: "Retrieve a list of the most recent jobs run.",
	Long:  `Retrieve a list of the most recent jobs run.`,
	Run: func(cmd *cobra.Command, args []string) {
		jsonString, err := GetJobs(fmt.Sprint(maxJobs))
		if err != nil {
			fmt.Printf("%+v\n", err)
			os.Exit(1)
		}
		if len(OutFormat) == 0 {
			fmt.Printf("%s\n\n", jsonString)
		} else {
			printHeader := true
			err := OPrintFormatted(OutFormat, jsonString, printHeader)
			// for i := range jobsData {
			// 	err := OPrintFormatted(OutFormat, jobsData[i], printHeader)
			if err != nil {
				fmt.Printf("%+v\n", err)
				os.Exit(1)
			}
			// 	printHeader = false
			// }
		}
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(getjobsCmd)

	// Here you will define your flags and configuration settings.
	getjobsCmd.Flags().UintVarP(&maxJobs, "max", "m", 5, "Max number of jobs to return")
	//getjobsCmd.MarkFlagRequired("max")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getJobsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getJobsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// GetJobs Get details for [count] last jobs
func GetJobs(count string) (jsonString string, err error) {

	username := os.Getenv("SAUCE_USERNAME")
	accessKey := os.Getenv("SAUCE_ACCESS_KEY")

	client := &http.Client{}
	request, err := http.NewRequest("GET", apiURL+"/jobs?limit="+count+"&full=true", nil)
	request.SetBasicAuth(username, accessKey)
	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("the http request to get jobs failed with error %s\n", err)
		return jsonString, err
	}
	jsonBytes, errReading := ioutil.ReadAll(response.Body)
	if errReading != nil {
		return "", fmt.Errorf("error reading API response")
	}
	return string(jsonBytes), err

}

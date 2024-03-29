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
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var max uint

// getjoblogsCmd represents the getjoblogs command
var getjoblogsCmd = &cobra.Command{
	Use:   "joblogs -m {maxJobs} (defaults to 5)",
	Short: "Get sauce and selenium-server log file from recent jobs. Saves to ./saucedata/{jobID}",
	Long:  `Get sauce and selenium-server log file from recent jobs. Saves to ./saucedata/{jobID}`,
	Run: func(cmd *cobra.Command, args []string) {
		GetJobLogs(max)
	},
}

func init() {
	rootCmd.AddCommand(getjoblogsCmd)

	// Here you will define your flags and configuration settings.
	getjoblogsCmd.Flags().UintVarP(&max, "max", "m", 5, "Max number of jobs to return")
	getjoblogsCmd.MarkFlagRequired("max")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getjoblogsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getjoblogsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// GetJobLogs Gets job lobs for [count] last jobs
func GetJobLogs(max uint) {

	jsonString, err := GetJobs(fmt.Sprint(max))
	if err != nil {
		panic(err)
	}

	var resultArray []map[string]interface{} // result is going to be the array object created from the json string
	errToArray := json.Unmarshal([]byte(jsonString), &resultArray)
	if errToArray != nil {
		panic(errToArray)
	} else {
		for _, item := range resultArray {

			jobID := fmt.Sprintf("%v", item["id"])

			os.MkdirAll(("./saucedata/" + jobID), 0777)
			jobString := fmt.Sprintf("%+v", item)
			ioutil.WriteFile("./saucedata/"+jobID+"/"+jobID+"-job-object.json", []byte(jobString), 0777)

			jsonString, err := GetJobAssetList(jobID)
			if err != nil {
				fmt.Printf("error getting asset list: %s\n", err)
				continue
			}
			if strings.Contains(jsonString, "Job hasn't finished running") {
				continue
			}
			var assetList map[string]interface{} // result is going to be the array object created from the json string
			errToList := json.Unmarshal([]byte(jsonString), &assetList)
			if errToList != nil {
				fmt.Printf("error converting asset list: %s\n", err)
				continue
			}
			sauceLog, err := GetAssetFile(jobID, fmt.Sprintf("%v", assetList["sauce-log"]))
			if err == nil {
				err = ioutil.WriteFile("./saucedata/"+jobID+"/"+jobID+"-sauce-log.json", []byte(sauceLog), 0777)
			} else {
				fmt.Printf("sauce-log retrieval error: %s\n", err)
			}

			seleniumServerLog, err := GetAssetFile(jobID, fmt.Sprintf("%v", assetList["selenium-log"]))
			if err == nil {
				err = ioutil.WriteFile("./saucedata/"+jobID+"/"+jobID+"-selenium-server.log", []byte(seleniumServerLog), 0777)
			} else {
				fmt.Printf("selenium-server retrieval error: %s\n", err)
			}
		}
	}
}

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
	"os"
	"time"

	"github.com/spf13/cobra"
)

var max uint

// getjoblogsCmd represents the getjoblogs command
var getjoblogsCmd = &cobra.Command{
	Use:   "getjoblogs -m {maxJobs}",
	Short: "Get sauce and selenium-server log file from recent jobs. Saves to ./saucedata/{jobID}",
	Long:  `Get sauce and selenium-server log file from recent jobs. Saves to ./saucedata/{jobID}`,
	Run: func(cmd *cobra.Command, args []string) {
		GetJobLogs(max)
	},
}

func init() {
	rootCmd.AddCommand(getjoblogsCmd)

	// Here you will define your flags and configuration settings.
	getjoblogsCmd.Flags().UintVarP(&max, "max", "m", 1, "Max number of jobs to return")
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

	jobs, _, err := GetJobs(fmt.Sprint(max))

	if err != nil {
		fmt.Printf("ERROR: GetJobLogs received the error message  \"%s\"\n", err)
	} else {
		for _, job := range jobs {
			startTime := time.Unix(job.StartTime, 0)
			fmt.Printf("Start Time: %s, ", startTime)
			if job.Passed {
				fmt.Printf("PASSED \n")
			} else if len(job.Error) > 0 {
				fmt.Printf("FAILED, Saucelabs Error: %s\n", job.Error)
			} else {
				fmt.Printf("FAILED \n")
			}
			os.MkdirAll(("./saucedata/" + job.ID), 0777)
			jobString := fmt.Sprintf("%+v", job)
			ioutil.WriteFile("./saucedata/"+job.ID+"/"+job.ID+"-job-object.json", []byte(jobString), 0777)

			// TODO: if "Job hasn't finished running" then we should skip, it's currently written as the log for some reason
			// TODO: does catching an error for getjobassetlist handle this?
			assetList, _, err := GetJobAssetList(job.ID)
			if err != nil {
				fmt.Printf("error getting asset list: %s\n", err)
				continue
			}

			sauceLog, err := GetAssetFile(job.ID, assetList.SauceLog)
			if err == nil {
				err = ioutil.WriteFile("./saucedata/"+job.ID+"/"+job.ID+"-sauce-log.json", []byte(sauceLog), 0777)
			} else {
				fmt.Printf("sauce-log retrieval error: %s\n", err)
			}

			seleniumServerLog, err := GetAssetFile(job.ID, assetList.SeleniumLog)
			if err == nil {
				err = ioutil.WriteFile("./saucedata/"+job.ID+"/"+job.ID+"-selenium-server.log", []byte(seleniumServerLog), 0777)
			} else {
				fmt.Printf("selenium-server retrieval error: %s\n", err)
			}

		}
	}
}

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

// getjoblogsCmd represents the getjoblogs command
var getjoblogsCmd = &cobra.Command{
	Use:   "getjoblogs 5",
	Short: "Get sauce and selenium-server log file from recent jobs. Saves to ./saucedata/{jobID}",
	Long:  `TODO: long description -> Get sauce and selenium-server log file from recent jobs. Saves files to ./saucedata/{jobID}`,
	Run: func(cmd *cobra.Command, args []string) {
		var maxJobs string
		if len(args) == 1 {
			maxJobs = args[0]
		} else {
			fmt.Printf("upload requires a parameter to specify how many job logs to list\ntry the --help option\n")
			os.Exit(1)
		}
		GetJobLogs(maxJobs)
	},
}

func init() {
	rootCmd.AddCommand(getjoblogsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getjoblogsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getjoblogsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// GetJobLogs Gets job lobs for [count] last jobs
func GetJobLogs(count string) {

	jobs, err := GetJobs(count)

	if err != nil {
		fmt.Printf("ERROR: GetJobLogs received the error message  \"%s\"\n", err)
	} else {
		for _, job := range jobs {
			//startTimeString, err := strconv.ParseInt(job.StartTime, 10, 64)
			// if err != nil {
			// 	fmt.Printf("error converting timestamp to string: %s", err)
			// }
			startTime := time.Unix(job.StartTime, 0)
			fmt.Printf("Start Time: %s, ", startTime)
			if job.Passed {
				fmt.Printf("PASSED ")
			} else if len(job.Error) > 0 {
				fmt.Printf("FAILED SAUCELABS ERROR: %s, ", job.Error)
			} else {
				fmt.Printf("FAILED ")
			}
			fmt.Printf("Scenario: %s, ID: %s\n", job.Name, job.ID)
			os.MkdirAll(("./saucedata/" + job.ID), 0777)
			jobString := fmt.Sprintf("%+v", job)
			ioutil.WriteFile("./saucedata/"+job.ID+"/"+job.ID+"-job-object.txt", []byte(jobString), 0777)

			// TODO: if "Job hasn't finished running" then we should skip, it's currently written as the log for some reason
			// TODO: does catching an error for getjobassetlist handle this?
			assetList, err := GetJobAssetList(job.ID)
			if err == nil {
				fmt.Printf("saucelog: %s\n\n", assetList.SauceLog)
			} else {
				fmt.Printf("%s\n\n", err)
				continue
			}

			sauceLog, err := GetAssetFile(job.ID, assetList.SauceLog)
			if err == nil {
				err = ioutil.WriteFile("./saucedata/"+job.ID+"/"+job.ID+"-sauce-log.json", []byte(sauceLog), 0777)
			} else {
				fmt.Printf("%s\n", err)
			}

			seleniumServerLog, err := GetAssetFile(job.ID, assetList.SeleniumLog)
			if err == nil {
				err = ioutil.WriteFile("./saucedata/"+job.ID+"/"+job.ID+"-selenium-server.log", []byte(seleniumServerLog), 0777)
			} else {
				fmt.Printf("%s\n", err)
			}
		}
	}
}

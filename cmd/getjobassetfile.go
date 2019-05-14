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

var jobIDForFile string
var filename string

// getjobassetfileCmd represents the getjobassetfile command
var getjobassetfileCmd = &cobra.Command{
	Use:   "getjobassetfile -j {jobid} -f {filename}",
	Short: "Dowload a specific asset file.",
	Long:  `Dowload a specific asset file.`,
	Run: func(cmd *cobra.Command, args []string) {
		var fileContents, err = GetAssetFile(jobIDForFile, filename)
		if err != nil {
			fmt.Printf("%s\n", err)
		}
		fmt.Printf("%s\n", fileContents)
	},
}

func init() {
	rootCmd.AddCommand(getjobassetfileCmd)

	getjobassetfileCmd.Flags().StringVarP(&jobIDForFile, "jobid", "j", "", "Saucelabs job ID")
	getjobassetfileCmd.MarkFlagRequired("jobid")
	getjobassetfileCmd.Flags().StringVarP(&filename, "filename", "f", "", "Name of the jobs asset file (see output from 'getjobassetfilelist')")
	getjobassetfileCmd.MarkFlagRequired("filename")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getjobassetfileCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getjobassetfileCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// GetAssetFile gets a copy of the specified file associate tot he job (screenshots, logs, etc)
func GetAssetFile(jobID string, filename string) (fileContents string, err error) {

	username := os.Getenv("SAUCE_USERNAME")
	accessKey := os.Getenv("SAUCE_ACCESS_KEY")

	client := &http.Client{}

	var url = apiURL + "/" + username + "/jobs/" + jobID + "/assets/" + filename
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("Error _creating request object_ to get asset file %s for job %s\n%s\n", filename, jobID, err)
	} else {
		request.SetBasicAuth(username, accessKey)
		response, err := client.Do(request)
		if err != nil {
			fmt.Printf("Error getting %s for job %s\n%s\n", filename, jobID, err)
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			fileContents = string(data)
		}
	}
	return fileContents, err
}

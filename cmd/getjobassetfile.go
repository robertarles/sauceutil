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

// getjobassetfileCmd represents the getjobassetfile command
var getjobassetfileCmd = &cobra.Command{
	Use:   "getjobassetfile {jobid} {filename}",
	Short: "Dowload a specific asset file.",
	Long:  `TODO: longer desc`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("getjobassetfile called")
	},
}

func init() {
	rootCmd.AddCommand(getjobassetfileCmd)

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

	// out, err := os.Create(jobID + "-" + filename)
	// defer out.Close()
	request, err := http.NewRequest("GET", apiURL+"/"+username+"/jobs/"+jobID+"/assets/"+filename, nil)
	if err != nil {
		fmt.Printf("Error _creating request object_ to get asset file %s for job %s\n%s", filename, jobID, err)
	} else {
		request.SetBasicAuth(username, accessKey)
		response, err := client.Do(request)
		//n, err := io.Copy(out, response.Body) copy to file
		if err != nil {
			fmt.Printf("Error getting %s for job %s\n%s\n", filename, jobID, err)
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			//fmt.Printf("Body of %s\n%s", filename, string(data))
			//defer response.Body.Close()
			// fmt.Printf("DEBUG RESPONSE N %v", n)
			// fmt.Printf("DEBUG response: \n%v\n", response)
			fileContents = string(data)
		}
	}
	return fileContents, err
}

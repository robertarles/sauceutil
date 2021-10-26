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

package sauceAPI

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// GetAssetFile gets a copy of the specified file associate tot he job (screenshots, logs, etc)
func GetAssetFile(jobID string, filename string) (fileContents string, err error) {

	// TODO: currently sending file streams to stdout, so must redirect to file
	// Send binaries direct to file rather than stdout? Why not the same for text files?

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

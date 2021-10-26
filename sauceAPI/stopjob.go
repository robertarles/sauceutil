// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
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
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

// StopJob Get detail on the specific job ID
// Note that a stop job request does not return a json response, only an HTTP status code
func StopJob(jobID string) (statusCode int, err error) {

	// TODO: add check for jobs existence and running status, then after stop command -> test if it's actually stopped

	username := os.Getenv("SAUCE_USERNAME")
	accessKey := os.Getenv("SAUCE_ACCESS_KEY")

	jobJson, getJobErr := GetJob(jobID)
	if getJobErr != nil {
		return 0, getJobErr
	}
	// convert the job response to a JobData struct
	var jobData JobData
	errUnmarshalling := json.Unmarshal([]byte(jobJson), &jobData)
	if errUnmarshalling != nil {
		return 0, errUnmarshalling
	}
	if (strings.ToLower(jobData.Status) != "running") && (strings.ToLower(jobData.Status) != "in progress") {
		return 0, fmt.Errorf(fmt.Sprintf("job %s does not appear to be running\n", jobID))
	}

	client := &http.Client{}
	request, reqErr := http.NewRequest("PUT", apiURL+"/"+username+"/jobs/"+jobID+"/stop", nil)
	if reqErr != nil {
		return 0, reqErr
	}
	request.SetBasicAuth(username, accessKey)
	response, doErr := client.Do(request)
	if doErr != nil {
		return 0, doErr
	}
	statusCode = response.StatusCode
	if statusCode > 299 {
		return 0, fmt.Errorf(fmt.Sprintf("stop job request returned status code %d\n", statusCode))
	}

	return statusCode, nil

}

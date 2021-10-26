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

// DeleteJob Get detail on the specific job ID
func DeleteJob(deleteJobID string) (deleteJSONString string, err error) {

	username := os.Getenv("SAUCE_USERNAME")
	accessKey := os.Getenv("SAUCE_ACCESS_KEY")

	jobJson, getErr := GetJob(deleteJobID)

	if getErr != nil {
		return "{}", getErr
	}
	// convert the job response to a JobData struct
	var jobData JobData
	errUnmarshalling := json.Unmarshal([]byte(jobJson), &jobData)
	if errUnmarshalling != nil {
		return "{}", errUnmarshalling
	}
	if jobData.ID != deleteJobID {
		return "", fmt.Errorf(fmt.Sprintf("job %s not found", deleteJobID))
	}
	if strings.ToLower(jobData.Status) == "running" {
		return "", fmt.Errorf(fmt.Sprintf("job %s is running", deleteJobID))
	}

	client := &http.Client{}
	request, reqErr := http.NewRequest("DELETE", apiURL+"/"+username+"/jobs/"+deleteJobID, nil)
	if reqErr != nil {
		return "", reqErr
	}
	request.SetBasicAuth(username, accessKey)
	response, doErr := client.Do(request)
	if doErr != nil {
		return "", err
	}
	if response.StatusCode > 299 {
		return "", fmt.Errorf(fmt.Sprintf("request error with status code %d\n", response.StatusCode))
	}

	respBody := DeleteJobData{}
	decoder := json.NewDecoder(response.Body)
	decodeErr := decoder.Decode(&respBody)
	if decodeErr != nil {
		return "", decodeErr
	}
	jsonBytes, marshErr := json.MarshalIndent(respBody, "", "  ")
	if marshErr != nil {
		return "", marshErr
	}
	return string(jsonBytes), nil

}

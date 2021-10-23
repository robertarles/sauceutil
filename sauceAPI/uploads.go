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
	"io/ioutil"
	"net/http"
	"os"
)

// Uploads returns a list of files in the sauce-storage upload area
func Uploads() (jsonString string, err error) {

	username := os.Getenv("SAUCE_USERNAME")
	accessKey := os.Getenv("SAUCE_ACCESS_KEY")

	// create the request object
	request, reqErr := http.NewRequest("GET", apiURL+"/storage/"+username, nil)
	if reqErr != nil {
		return "", reqErr
	}

	request.SetBasicAuth(username, accessKey)
	client := &http.Client{}
	response, doErr := client.Do(request)
	if doErr != nil {
		return "", doErr
	}
	jsonBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(jsonBytes), nil

}

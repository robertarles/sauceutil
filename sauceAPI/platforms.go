//https://api.us-west-1.saucelabs.com/rest/v1/info/platforms/all
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
)

// Get platforms available for testing
func Platforms(platform string) (jsonString string, err error) {

	// make the request
	response, err := http.Get(apiURL + "/info/platforms/" + platform)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return "", err
	}

	// decode the response body
	body, errReading := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if errReading != nil {
		return "", fmt.Errorf("error reading API response")
	}
	return string(body), nil

}

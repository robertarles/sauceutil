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
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// Tunnel Get detail on the specific job ID
func Tunnel(tunnelID string) (jsonString string, err error) {

	username := os.Getenv("SAUCE_USERNAME")
	accessKey := os.Getenv("SAUCE_ACCESS_KEY")

	client := &http.Client{}
	request, err := http.NewRequest("GET", apiURL+"/"+username+"/tunnels/"+tunnelID, nil)
	request.SetBasicAuth(username, accessKey)
	response, err := client.Do(request)
	if err != nil {
		return fmt.Sprintf(`{"error": "The http request failed with error %s}"`, err), err
	}

	respBody := TunnelData{}
	decoder := json.NewDecoder(response.Body)
	fmt.Printf("[TODO fix handling of tunnel data] [DEBUG] %+v\n", respBody)
	os.Exit(0)
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

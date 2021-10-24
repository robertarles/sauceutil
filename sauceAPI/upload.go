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
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

// Upload Post a file to sauce-storage
func Upload(uploadFilepath string) (jsonString string, err error) {

	username := os.Getenv("SAUCE_USERNAME")
	accessKey := os.Getenv("SAUCE_ACCESS_KEY")

	uploadFilename := filepath.Base(uploadFilepath)
	//open file and retrieve info
	file, err := os.Open(uploadFilepath)
	if err != nil {
		fmt.Println(err)
	}
	//fileContents, err := ioutil.ReadAll(file)
	body := &bytes.Buffer{}
	io.Copy(body, file)

	postURL := apiURL + "/storage/" + username + "/" + uploadFilename + "?overwrite=true"

	request, err := http.NewRequest("POST", postURL, body)
	request.SetBasicAuth(username, accessKey)
	request.Header.Add("Content-Type", "application/octet-stream")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return "", err
	} else if response.StatusCode != 200 {
		return fmt.Sprintf(`"message": "non-200 http response", "status_code": "%d"}`, response.StatusCode), nil
	}

	respBody := UploadResponse{}
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

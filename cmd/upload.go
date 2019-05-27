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
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var uploadFilename string

// uploadCmd represents the upload command
var uploadCmd = &cobra.Command{
	Use:   "upload -f {filename}",
	Short: "Upload a file to your sauce-storage temp file storage area.",
	Long:  `Upload a file to your sauce-storage temp file storage area.`,
	Run: func(cmd *cobra.Command, args []string) {
		var uploadResponse, jsonString, err = Upload(uploadFilename)
		if err != nil {
			fmt.Printf("%s\n", err)
			os.Exit(1)
		}
		if err != nil {
			fmt.Printf("%+v\n", err)
			os.Exit(1)
		}
		if len(OutFormat) == 0 {
			fmt.Printf("%s\n", jsonString)
		} else {
			printHeader := true
			err := OPrintStruct(OutFormat, uploadResponse, printHeader)
			if err != nil {
				fmt.Printf("%+v\n", err)
				os.Exit(1)
			}
		}
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(uploadCmd)
	// Here you will define your flags and configuration settings.
	uploadCmd.Flags().StringVarP(&uploadFilename, "filename", "f", "", "Name of file to upload to sauce-storage")
	uploadCmd.MarkFlagRequired("filename")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// uploadCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// uploadCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// Upload Post a file to sauce-storage
func Upload(uploadFilepath string) (uploadResponseData UploadResponse, jsonString string, err error) {

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
		return UploadResponse{}, "", err
	} else if response.StatusCode != 200 {
		return UploadResponse{}, fmt.Sprintf(`"message": "non-200 http response", "status_code": "%d"}`, response.StatusCode), nil
	}

	respBody := UploadResponse{}
	decoder := json.NewDecoder(response.Body)
	decodeErr := decoder.Decode(&respBody)
	if decodeErr != nil {
		return UploadResponse{}, "", decodeErr
	}
	jsonBytes, marshErr := json.MarshalIndent(respBody, "", "  ")
	if marshErr != nil {
		return UploadResponse{}, "", marshErr
	}
	return respBody, string(jsonBytes), nil

}

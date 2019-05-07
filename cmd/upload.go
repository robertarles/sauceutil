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

package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// uploadCmd represents the upload command
var uploadCmd = &cobra.Command{
	Use:   "upload {filename}",
	Short: "Upload a file to your sauce-storage temp file storage area.",
	Long:  `TODO: longer desc -> Upload a file to your sauce-storage temp file storage area.`,
	Run: func(cmd *cobra.Command, args []string) {
		var filename string
		if len(args) == 1 {
			filename = args[0]
		} else {
			fmt.Printf("upload requires a filename parameter\ntry the --help option\n")
			os.Exit(1)
		}
		Upload(filename)
	},
}

func init() {
	rootCmd.AddCommand(uploadCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// uploadCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// uploadCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type uploadResponse struct {
	Username string `json:"username"`
	Filename string `json:"filename"`
	Size     string `json:"size"`
	Md5      string `json:"md5"`
	Etag     string `json:"etag"`
}

// PostUpload Post a file to sauce-storage
func Upload(uploadFilepath string) {

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

	fmt.Printf("Posting to %s\n", postURL)
	request, err := http.NewRequest("POST", postURL, body)
	request.SetBasicAuth(username, accessKey)
	request.Header.Add("Content-Type", "application/octet-stream")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("The http request failed with error %s\n", err)
	} else if response.StatusCode != 200 {
		fmt.Printf("Upload request failed with status code of %d\n", response.StatusCode)
	} else {
		respBody := uploadResponse{}
		data, _ := ioutil.ReadAll(response.Body)
		json.Unmarshal(data, &respBody)
		fmt.Printf("\nUploaded - Username: %s, File: %s, Size: %s, md5: %s", respBody.Username, respBody.Filename, respBody.Size, respBody.Md5)
	}

}

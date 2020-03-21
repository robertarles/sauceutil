package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var tunnelsCmd = &cobra.Command{
	Use:   "tunnels",
	Short: "A list of tunnels available to your account.",
	Long:  `A list of tunnels available to your account.`,
	Run: func(cmd *cobra.Command, args []string) {
		var tunnelsResponse, jsonString, err = Tunnels()
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
			err := OPrintStruct(OutFormat, tunnelsResponse, printHeader)
			if err != nil {
				fmt.Printf("%+v\n", err)
				os.Exit(1)
			}
		}
		os.Exit(0)

	},
}

func init() {
	rootCmd.AddCommand(tunnelsCmd)
}

// Tunnels returns a list of tunnels available to your account
func Tunnels() (tunnelsResponse []string, jsonString string, err error) {

	username := os.Getenv("SAUCE_USERNAME")
	accessKey := os.Getenv("SAUCE_ACCESS_KEY")

	client := &http.Client{}
	request, reqErr := http.NewRequest("GET", apiURL+"/"+username+"/tunnels", nil)
	if reqErr != nil {
		return []string{}, "", reqErr
	}
	request.SetBasicAuth(username, accessKey)
	response, doErr := client.Do(request)
	if doErr != nil {
		return []string{}, "", doErr
	}

	decoder := json.NewDecoder(response.Body)
	tunnelsResponse = []string{}
	decodeErr := decoder.Decode(&tunnelsResponse)
	if decodeErr != nil {
		return []string{}, "", decodeErr
	}

	if len(tunnelsResponse) < 1 {
		return []string{}, "", nil
	}

	jsonBytes, marshErr := json.MarshalIndent(tunnelsResponse, "", "  ")
	if marshErr != nil {
		return []string{}, "", marshErr
	}
	return tunnelsResponse, string(jsonBytes), nil

}

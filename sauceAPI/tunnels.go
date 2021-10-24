package sauceAPI

import (
	"encoding/json"
	"net/http"
	"os"
)

// Tunnels returns a list of tunnels available to your account
func Tunnels() (jsonString string, err error) {

	username := os.Getenv("SAUCE_USERNAME")
	accessKey := os.Getenv("SAUCE_ACCESS_KEY")

	client := &http.Client{}
	request, reqErr := http.NewRequest("GET", apiURL+"/"+username+"/tunnels", nil)
	if reqErr != nil {
		return "", reqErr
	}
	request.SetBasicAuth(username, accessKey)
	response, doErr := client.Do(request)
	if doErr != nil {
		return "", doErr
	}

	decoder := json.NewDecoder(response.Body)
	tunnelsResponse := []string{}
	decodeErr := decoder.Decode(&tunnelsResponse)
	if decodeErr != nil {
		return "", decodeErr
	}

	if len(tunnelsResponse) < 1 {
		return "", nil
	}

	jsonBytes, marshErr := json.MarshalIndent(tunnelsResponse, "", "  ")
	if marshErr != nil {
		return "", marshErr
	}
	return string(jsonBytes), nil

}

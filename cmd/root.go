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
	"fmt"
	"os"

	sauceAPI "github.com/robertarles/sauceutil/sauceAPI"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// OutFormat an array of user supplied fields to determine columns displayed for formatted output
var OutFormat []string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sauceutil",
	Short: "A command line utility for Saucelabs API tasks.",
	Long:  "A command line utility for Saucelabs tasks.\nEasily upload, check uploads, get job assets and info from the command line.",
}

// apistatusCmd represents the apistatus command
var apistatusCmd = &cobra.Command{
	Use:   "apistatus",
	Short: "Request the current API status.",
	Long:  `Request the current API status.`,
	Run: func(cmd *cobra.Command, args []string) {
		var jsonString, err = sauceAPI.GetAPIStatus()
		sauceAPI.Log(jsonString, err, OutFormat)
	},
}

// uploadsCmd represents the uploads command
var uploadsCmd = &cobra.Command{
	Use:   "uploads",
	Short: "A list of files already uploaded to sauce-storage.",
	Long:  `A list of files already uploaded to sauce-storage.`,
	Run: func(cmd *cobra.Command, args []string) {
		var jsonString, err = sauceAPI.Uploads()
		sauceAPI.Log(jsonString, err, OutFormat)
	},
}

var uploadFilename string

// uploadCmd represents the upload command
var uploadCmd = &cobra.Command{
	Use:   "upload -f {filename}",
	Short: "Upload a file to your sauce-storage temp file storage area.",
	Long:  `Upload a file to your sauce-storage temp file storage area.`,
	Run: func(cmd *cobra.Command, args []string) {
		var jsonString, err = sauceAPI.Upload(uploadFilename)
		sauceAPI.Log(jsonString, err, OutFormat)
	},
}

var tunnelsCmd = &cobra.Command{
	Use:   "tunnels",
	Short: "A list of tunnels available to your account.",
	Long:  `A list of tunnels available to your account.`,
	Run: func(cmd *cobra.Command, args []string) {
		var jsonString, err = sauceAPI.Tunnels()
		sauceAPI.Log(jsonString, err, OutFormat)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringSliceVarP(&OutFormat, "", "o", []string{}, "Formatted output. Supply a single, quoted and comma separated list of columns to display")
	rootCmd.AddCommand(apistatusCmd)
	rootCmd.AddCommand(uploadsCmd)
	rootCmd.AddCommand(tunnelsCmd)
	rootCmd.AddCommand(uploadCmd)
	// define your flags and configuration settings.
	uploadCmd.Flags().StringVarP(&uploadFilename, "filename", "f", "", "Name of file to upload to sauce-storage")
	uploadCmd.MarkFlagRequired("filename")

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".sauceutil" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".sauceutil")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

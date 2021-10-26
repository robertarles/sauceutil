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
		sauceAPI.LogJsonResults(jsonString, err, OutFormat)
	},
}

// uploadsCmd represents the uploads command
var uploadsCmd = &cobra.Command{
	Use:   "uploads",
	Short: "A list of files already uploaded to sauce-storage.",
	Long:  `A list of files already uploaded to sauce-storage.`,
	Run: func(cmd *cobra.Command, args []string) {
		var jsonString, err = sauceAPI.Uploads()
		sauceAPI.LogJsonResults(jsonString, err, OutFormat)
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
		sauceAPI.LogJsonResults(jsonString, err, OutFormat)
	},
}

var tunnelsCmd = &cobra.Command{
	Use:   "tunnels",
	Short: "A list of tunnels available to your account.",
	Long:  `A list of tunnels available to your account.`,
	Run: func(cmd *cobra.Command, args []string) {
		var jsonString, err = sauceAPI.Tunnels()
		sauceAPI.LogJsonResults(jsonString, err, OutFormat)
	},
}

var maxJobs uint

// getJobsCmd represents the getJobs command
var getjobsCmd = &cobra.Command{
	Use:   "jobs -m {maxJobs} (defaults to 5)",
	Short: "Retrieve a list of the most recent jobs run.",
	Long:  `Retrieve a list of the most recent jobs run.`,
	Run: func(cmd *cobra.Command, args []string) {
		var jsonString, err = sauceAPI.GetJobs(fmt.Sprint(maxJobs))
		sauceAPI.LogJsonResults(jsonString, err, OutFormat)
	},
}

var max uint

// getjoblogsCmd represents the getjoblogs command
var getjoblogsCmd = &cobra.Command{
	Use:   "joblogs -m {maxJobs} (defaults to 5)",
	Short: "Get sauce and selenium-server log file from recent jobs. Saves to ./saucedata/{jobID}",
	Long:  `Get sauce and selenium-server log file from recent jobs. Saves to ./saucedata/{jobID}`,
	Run: func(cmd *cobra.Command, args []string) {
		sauceAPI.GetJobLogs(max)
	},
}

var getJobID string

// getJobCmd represents the getJob command
var getjobCmd = &cobra.Command{
	Use:   "job -i {jobID}",
	Short: "Get details on a specific job",
	Long:  `Get details on a specific job`,
	Run: func(cmd *cobra.Command, args []string) {
		var jsonString, err = sauceAPI.GetJob(getJobID)
		sauceAPI.LogJsonResults(jsonString, err, OutFormat)
	},
}

var deleteJobID string

// deletejobCmd represents the deletejob command
var deletejobCmd = &cobra.Command{
	Use:   "deletejob",
	Short: "Removes the job from the Saucelabs system with all the linked assets",
	Long:  `Removes the job from the Saucelabs system with all the linked assets`,
	Run: func(cmd *cobra.Command, args []string) {
		var jsonString, err = sauceAPI.DeleteJob(deleteJobID)
		sauceAPI.LogJsonResults(jsonString, err, OutFormat)
	},
}
var stopJobID string

// stopjobCmd represents the stopjob command
var stopjobCmd = &cobra.Command{
	Use:   "stopjob -i {jobID}",
	Short: "Terminates a running Saucelabs job",
	Long:  `Terminates a running Saucelabs job`,
	Run: func(cmd *cobra.Command, args []string) {
		var statusCode, err = sauceAPI.StopJob(stopJobID)
		// special case logging, no json, just status code
		if err != nil {
			fmt.Printf("%s\n", err)
			os.Exit(1)
		}
		fmt.Printf("http %d\n", statusCode)
	},
}

var jobIDForFile string
var filename string

// getjobassetfileCmd represents the getjobassetfile command
var getjobassetfileCmd = &cobra.Command{
	Use:   "assetfile -i {jobID} -f {filename}",
	Short: "Dowload a specific asset file.",
	Long:  `Dowload a specific asset file.`,
	Run: func(cmd *cobra.Command, args []string) {
		var fileContents, err = sauceAPI.GetAssetFile(jobIDForFile, filename)
		// special case, dump file text, no json to process
		if err != nil {
			fmt.Printf("%s\n", err)
			os.Exit(1)
		}
		fmt.Printf("%s\n", fileContents)
	},
}

var jobID string

// getjobassetlistCmd represents the getjobassetlist command
var getjobassetlistCmd = &cobra.Command{
	Use:   "assetlist -i {JobID}",
	Short: "Get a list of files associated to a job.",
	Long:  `Get a list of files associated to a job.`,
	Run: func(cmd *cobra.Command, args []string) {
		var jsonString, err = sauceAPI.GetJobAssetList(jobID)
		sauceAPI.LogJsonResults(jsonString, err, OutFormat)
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

// configure the CLI handling
func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringSliceVarP(&OutFormat, "", "o", []string{}, "Formatted output. Supply a single, quoted and comma separated list of columns to display")
	rootCmd.AddCommand(apistatusCmd)
	rootCmd.AddCommand(uploadsCmd)
	rootCmd.AddCommand(tunnelsCmd)
	rootCmd.AddCommand(uploadCmd)
	uploadCmd.Flags().StringVarP(&uploadFilename, "filename", "f", "", "Name of file to upload to sauce-storage")
	uploadCmd.MarkFlagRequired("filename")
	rootCmd.AddCommand(getjobsCmd)
	getjobsCmd.Flags().UintVarP(&maxJobs, "max", "m", 5, "Max number of jobs to return")
	rootCmd.AddCommand(getjoblogsCmd)
	getjoblogsCmd.Flags().UintVarP(&max, "max", "m", 5, "Max number of jobs to return")
	getjoblogsCmd.MarkFlagRequired("max")
	rootCmd.AddCommand(getjobCmd)
	getjobCmd.Flags().StringVarP(&getJobID, "id", "i", "", "Saucelabs Job ID")
	getjobCmd.MarkFlagRequired("id")
	rootCmd.AddCommand(deletejobCmd)
	deletejobCmd.Flags().StringVarP(&deleteJobID, "id", "i", "", "Saucelabs Job ID")
	deletejobCmd.MarkFlagRequired("id")
	rootCmd.AddCommand(stopjobCmd)
	stopjobCmd.Flags().StringVarP(&stopJobID, "id", "i", "", "Saucelabs Job ID")
	stopjobCmd.MarkFlagRequired("id")
	rootCmd.AddCommand(getjobassetfileCmd)
	getjobassetfileCmd.Flags().StringVarP(&jobIDForFile, "id", "i", "", "Saucelabs job ID")
	getjobassetfileCmd.MarkFlagRequired("id")
	getjobassetfileCmd.Flags().StringVarP(&filename, "filename", "f", "", "Name of the jobs asset file (see output from 'getjobassetfilelist')")
	getjobassetfileCmd.MarkFlagRequired("filename")
	rootCmd.AddCommand(getjobassetlistCmd)
	getjobassetlistCmd.Flags().StringVarP(&jobID, "id", "i", "", "The Saucelabs job ID to get an asset list for.")
	getjobassetlistCmd.MarkFlagRequired("id")
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

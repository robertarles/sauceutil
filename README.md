# Sauceutil

A Saucelabs command line utility written in Go.

A quick way to get things done from the command line, more portable than an npm package since the executable can be shared once compiled, but for a more complete util, check out Saucelabs own https://www.npmjs.com/package/saucelabs.

## Installation

`go get github.com/robertarles/sauceutil`

`sauceutil` expects your username and access key to be set in your environment variables (eg ~/.bashrc or ~/.zshrc):

``` bash
export SAUCE_USERNAME=yourAccountUsername
export SAUCE_ACCESS_KEY=yourAPIAccessKey
```

## Use

``` text
A command line utility for Saucelabs tasks.
Easily upload, check uploads, get job assets and info from the command line.

Usage:
  sauceutil [command]

Available Commands:
  apistatus       Request the current API status.
  getjob          Get details on a specific job
  getjobassetfile Dowload a specific asset file.
  getjobassetlist Get a list of files associated to a job.
  getjoblogs      Get sauce and selenium-server log file from recent jobs. Saves to ./saucedata/{jobID}
  getjobs         Retrieve a list of the most recent jobs run.
  help            Help about any command
  upload          Upload a file to your sauce-storage temp file storage area.
  uploads         A list of files already uploaded to sauce-storage.

Flags:
  -h, --help   help for sauceutil

Use "sauceutil [command] --help" for more information about a command.

```
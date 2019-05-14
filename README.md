# Sauceutil

A Saucelabs command line utility written in Go.

A quick way to get things done from the command line, more portable than an npm package since the executable can be shared once compiled, it's fast, it's small, and no need to install node+npm; but for a more complete util, check out Saucelabs own https://www.npmjs.com/package/saucelabs.

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

``` bash
# sauceutil uploads  
{
  "files": [
    {
      "name": "Android_App.apk",
      "size": 12194252,
      "mtime": 1557162500,
      "md5": "daf275a1bd0e4672023f4c6d38a03063",
      "etag": "1fe0092eae16346c75132f50e73e7b7e"
    },
    {
      "name": "iOS_App.zip",
      "size": 22529112,
      "mtime": 1557163100,
      "md5": "660a6591285b94a433a85914b9512056b",
      "etag": "5a352249cf71f433b3b8060465d2a5b9"
    }
  ]
}  
```

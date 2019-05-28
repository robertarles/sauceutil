# Sauceutil

A Saucelabs command line utility written in Go.

A quick way to get things done from the command line, more portable than an npm package since the executable can be shared once compiled, it's fast, it's small, and no need to install node+npm; but for a more complete util, check out Saucelabs own https://www.npmjs.com/package/saucelabs.

## Installation

Just download the binary for your system from the `bin` directory. Navigate to the executable above, click it, then you'll see a download button.

-OR-

`go get github.com/robertarles/sauceutil`

## Configuration

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
  deletejob       Removes the job from the Saucelabs system with all the linked assets
  getjob          Get details on a specific job
  getjobassetfile Dowload a specific asset file.
  getjobassetlist Get a list of files associated to a job.
  getjoblogs      Get sauce and selenium-server log file from recent jobs. Saves to ./saucedata/{jobID}
  getjobs         Retrieve a list of the most recent jobs run.
  help            Help about any command
  stopjob         Terminates a running Saucelabs job
  upload          Upload a file to your sauce-storage temp file storage area.
  uploads         A list of files already uploaded to sauce-storage.

Flags:
  -o, -- strings   Formatted output. Supply a single, quoted and comma separated list of columns to display
  -h, --help       help for sauceutil

Use "sauceutil [command] --help" for more information about a command.

```

### Command line example

``` bash
$ sauceutil getjobs -m 1
[
  {
    "browser_short_version": "8.0",
    "video_url": "https://assets.saucelabs.com/jobs/3ec4af43fa7a486b9217e733a798bc85/video.flv",
    "creation_time": 1557161283,
    "browser_version": "8.0.",
    "owner": "ownername",
    "id": "3ec4af43fa7a486b9217e733a798bc85",
    "container": false,
    "record_screenshots": true,
    "record_video": true,
    "build": "20190301_14.1",
    "passed": false,
    "public": "team",
    "end_time": 1557161283,
    "status": "complete",
    "log_url": "https://assets.saucelabs.com/jobs/3ec4af43fa7a486b9217e733a798bc85/selenium-server.log",
    "start_time": 0,
    "proxied": false,
    "modification_time": 1557161283,
    "name": "",
    "commands_not_successful": 0,
    "consolidated_stats": "",
    "assigned_tunnel_id": "",
    "error": "No active tunnel found for identifier primary_sauce_tunnel",
    "os": "Linux",
    "breakpointed": false,
    "browser": "android"
  }
]
  
```
### Command line example with `-o` formatting. Note the field names are from the JSON in the previous example
``` bash
$ sauceutil getjobs -m 5 -o "id,passed,status,owner"
id                                passed  status    owner              
3ec4af43fa7a486b9217e733a798bc85  false   complete  ownername1  
0f2833b8492b424ae604a90f556e915c  false   complete  ownername3  
54ea20cd3f6b4db8be0984e0d42a6376  false   complete  ownername2  
c029e776f1b847b580a27e0dd53d198e  false   complete  ownername1  
d56880d2b83af6f0a456417f127af80f  false   complete  ownername1 
```

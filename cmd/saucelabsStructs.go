package cmd

// FileData structure of Saucelabs API response
type FileData struct {
	Name  string  `json:"name"`
	Size  uint32  `json:"size"`
	Mtime float32 `json:"mtime"`
	Md5   string  `json:"md5"`
	Etag  string  `json:"etag"`
}

// StorageResponse structure of a Saucelabs API response
type StorageResponse struct {
	Files []FileData `json:"files"`
}

// CustomJobData internal structure of a Saucelabs JobData API response
type CustomJobData struct {
	BuildNumber      string `json:"BUILD_NUMBER"`
	JenkinsBuildName string `json:"JENKINS_BUILD_NAME"`
	GitCommit        string `json:"GIT_COMMIT"`
}

// JobData structure of a Saucelabs API response
type JobData struct {
	BrowserShortVersion string `json:"browser_short_version"`
	VideoURL            string `json:"video_url"`
	CreationTime        int64  `json:"creation_time"`
	//CustomData            CustomJobData `json:"custom-data"` // TODO: handle the saucelabs  error? sometimes returns CustomData.BuildNumber as an int, not string
	BrowserVersion        string   `json:"browser_version"`
	Owner                 string   `json:"owner"`
	ID                    string   `json:"id"`
	Container             bool     `json:"container"`
	RecordScreenshots     bool     `json:"record_screenshots"`
	RecordVideo           bool     `json:"record_video"`
	Build                 string   `json:"build"`
	Passed                bool     `json:"passed"`
	Public                string   `json:"public"`
	EndTime               int64    `json:"end_time"`
	Status                string   `json:"status"`
	LogURL                string   `json:"log_url"`
	StartTime             int64    `json:"start_time"`
	Proxied               bool     `json:"proxied"`
	ModificationTime      int64    `json:"modification_time"`
	Tags                  []string `json:"tags"`
	Name                  string   `json:"name"`
	CommandsNotSuccessful uint32   `json:"commands_not_successful"`
	ConsolidatedStatus    string   `json:"consolidated_stats"`
	AssignedTunnelID      string   `json:"assigned_tunnel_id"`
	Error                 string   `json:"error"`
	OS                    string   `json:"os"`
	Breakpointed          bool     `json:"breakpointed"`
	Browser               string   `json:"browser"`
}

// UploadResponse structure of a Saucelabs API response
type UploadResponse struct {
	Username string `json:"username"`
	Filename string `json:"filename"`
	Size     string `json:"size"`
	Md5      string `json:"md5"`
	Etag     string `json:"etag"`
}

// AssetListData structure of a Saucelabs API response
type AssetListData struct {
	SauceLog    string   `json:"sauce-log"`
	Video       string   `json:"video"`
	SeleniumLog string   `json:"selenium-log"`
	Screenshots []string `json:"screenshots"`
}

// APIStatusResponseData structure of a Saucelabs API response
type APIStatusResponseData struct {
	WaitTime           float64 `json:"wait_time"`
	ServiceOperational bool    `json:"service_operational"`
	StatusMessage      string  `json:"status_message"`
}

// DeleteJobData struct of a Saucelabs delete job response
type DeleteJobData struct {
	Status string
}

// StopJobData struct of a Saucelabs stop job response
type StopJobData struct {
	Status string
}

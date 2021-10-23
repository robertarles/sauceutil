package sauceAPI

// default api url
var apiURL = "https://saucelabs.com/rest/v1"

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

// TunnelExtra is sub-object in a Saucelabs tunnel API response
type TunnelExtra struct {
	InjectJobID bool   `json:"inject_job_id"`
	Backend     string `json:"backend"`
	MetricsHost string `json:"metrics_host"`
	MetricsPort int    `json:"metrics_port"`
}

// TunnelMetadata is a sub-object in a saucelabs tunnel API response
type TunnelMetadata struct {
	Hostname    string `json:"hostname"`
	GitVersion  string `json:"git_version"`
	Platform    string `json:"platform"`
	Command     string `json:"command"`
	Build       string `json:"build"`
	Release     string `json:"release"`
	NofileLimit int64  `json:"nofile_limit"`
}

// TunnelData is Saucelabs Tunnel API response
type TunnelData struct {
	TeamIds          []string       `json:"team_ids"`
	SSHPort          int            `json:"ssh_port"`
	CreationTime     int64          `json:"creation_time"`
	DomainNames      []string       `json:"domain_names"`
	Owner            string         `json:"owner"`
	UseKGP           bool           `json:"use_kgp"`
	ID               string         `json:"id"`
	ExtraInfo        TunnelExtra    `json:"extra_info"`
	DirectDomains    []string       `json:"direct_domains"`
	VMVersion        string         `json:"vm_version"`
	NoSslBumpDomains []string       `json:"no_ssl_bump_domains"`
	SharedTunnel     bool           `json:"shared_tunnel"`
	Metadata         TunnelMetadata `json:"metadata"`
	Status           string         `json:"status"`
	ShutdownTime     string         `json:"shutdown_time"`
	Host             string         `json:"host"`
	IPAddress        string         `json:"ip_address"`
	LastConnected    string         `json:"last_connected"`
	UserShutdown     string         `json:"user_shutdown"`
	UseCachingProxy  bool           `json:"use_caching_proxy"`
	LaunchTime       int64          `json:"launch_time"`
	NoProxyCaching   bool           `json:"no_proxy_caching"`
	TunnelIdentifier string         `json:"tunnel_identifier"`
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
	Size     int    `json:"size"`
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
	Status string `json:"status"`
}

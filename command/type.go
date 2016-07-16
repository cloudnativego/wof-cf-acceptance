package command

import "github.com/fatih/color"

const (
	cfCommand      = "cf"
	configString   = "config"
	spacemanPlugin = "spaceman"
)

var (
	white  = color.New(color.FgWhite)
	bold   = color.New(color.FgWhite, color.Bold)
	yellow = color.New(color.FgYellow)
	red    = color.New(color.FgRed, color.Bold)
	green  = color.New(color.FgGreen, color.Bold)
	blue   = color.New(color.FgCyan, color.Bold)
)

//DeploymentConfig contains the information necessary to (un)deploy
//World of Fluxcraft to Cloud Foundry.
type DeploymentConfig struct {
	SpacemanConfig string
	WOFApps        []WOFApp
}

//SpaceConfiguration holds the services from spaceman config.
type SpaceConfiguration struct {
	UserProvided []Service
	Brokered     []Service
}

//Service holds a service name provided by spaceman config.
type Service struct {
	Name string
}

//WOFApp contains the information necessary to deploy a given app
//to Cloud Foundry.
type WOFApp struct {
	ProjectURI  string
	Name        string
	Host        string
	Memory      string
	Instances   int
	DiskQuota   string
	Buildpack   string
	Command     string
	Domain      string
	NoHost      bool
	NoRoute     bool
	RandomRoute bool
	Env         map[string]interface{}
}

func (s *WOFApp) applyDefaultValues() {
	if len(s.Memory) == 0 {
		s.Memory = "32M"
	}
	if s.Instances == 0 {
		s.Instances = 1
	}
	if len(s.DiskQuota) == 0 {
		s.DiskQuota = "100M"
	}
}

func (s *WOFApp) isValid() (valid bool, issues []string) {
	valid = true
	issues = make([]string, 0)

	if len(s.ProjectURI) == 0 {
		issues = append(issues, "A ProjectURI must be specified")
	}

	if len(s.Name) == 0 {
		issues = append(issues, "A name must be specified")
	}

	if len(s.Host) == 0 {
		issues = append(issues, "A host must be specified")
	}

	if len(issues) > 0 {
		valid = false
	}

	return
}

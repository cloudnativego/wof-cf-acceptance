package command

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/urfave/cli"
)

//CmdDeploy -
func CmdDeploy(c *cli.Context) error {
	bold.Println("Starting Deploy...")

	if !execInstalled(cfCommand) {
		writeErrorAndExit(fmt.Sprintf("Error: `%s` command was not found. Ensure you have `%s` installed and that it is available on the PATH", cfCommand, cfCommand))
	}
	white.Println("INFO: cf executable found")

	if !pluginInstalled(spacemanPlugin) {
		writeErrorAndExit(fmt.Sprintf("Error: `%s` plugin was not found. Ensure that you install the `%s` plugin (github.com/pivotalservices/cf-spaceman)", spacemanPlugin, spacemanPlugin))
	}
	white.Println("INFO: cf spaceman plugin found")

	configPath := c.String(configString)
	if len(configPath) == 0 {
		writeErrorAndExit("A path to a deployment config must be specified")
	}

	config, err := readFromFile(configPath)
	if err != nil {
		writeErrorAndExit(fmt.Sprintf("The deployment config could not be found at: %s", configPath))
	}
	white.Println("INFO: reading deployment config...")

	dc, err := parseDeploymentConfig(config)
	if err != nil {
		writeErrorAndExit("Error parsing deployment config")
	}
	white.Println("INFO: ...done!")

	if len(dc.SpacemanConfig) == 0 {
		writeErrorAndExit("A spaceman configuration file must be specified")
	}

	white.Printf("INFO: Deploying services with cf-spaceman using: %s...\n\n", dc.SpacemanConfig)

	cmd := exec.Command("cf", "spaceman", dc.SpacemanConfig)
	cmd.Stdout = os.Stdout
	err = cmd.Run()
	if err != nil {
		writeErrorAndExit(fmt.Sprintf("Spaceman Error: %s", err))
	}
	white.Println("INFO: ...done!")

	white.Println("INFO: Deploying applications...")
	for _, app := range dc.WOFApps {
		app.applyDefaultValues()

		if ok, issues := app.isValid(); !ok {
			red.Println("APPLCIATION CONFIGURATION ERROR:")
			for _, i := range issues {
				red.Println(i)
			}
		} else {
			//TODO: Deploy App
			yellow.Printf("%+v\n", app)
		}
	}
	white.Println("INFO: ...done!")

	bold.Println("INFO: Deployment succeeded.")
	green.Print("\nSUCCESS\n\n")
	return nil
}

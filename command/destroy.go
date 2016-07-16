package command

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/urfave/cli"
)

//CmdDestroy -
func CmdDestroy(c *cli.Context) error {
	bold.Println("Starting Destroy...")

	if !execInstalled(cfCommand) {
		writeErrorAndExit(fmt.Sprintf("Error: `%s` command was not found. Ensure you have `%s` installed and that it is available on the PATH", cfCommand, cfCommand))
	}
	white.Println("INFO: cf executable found")

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

	//TODO: REMOVE APPS FIRST

	if len(dc.SpacemanConfig) == 0 {
		writeErrorAndExit("A spaceman configuration file must be specified")
	}

	spaceConfig, err := readFromFile(dc.SpacemanConfig)
	if err != nil {
		writeErrorAndExit(fmt.Sprintf("The spaceman config could not be found at: %s", dc.SpacemanConfig))
	}

	white.Println("INFO: reading spaceman config...")
	sc, err := parseSpaceConfiguration(spaceConfig)
	if err != nil {
		writeErrorAndExit("Error parsing spaceman config")
	}
	white.Println("INFO: ...done")

	var cmd *exec.Cmd

	//refactor
	white.Println("INFO: Deleting user provided services...")
	for _, ups := range sc.UserProvided {
		cmd = exec.Command("cf", "ds", ups.Name, "-f")
		cmd.Stdout = os.Stdout
		err = cmd.Run()
		if err != nil {
			writeErrorAndExit(fmt.Sprintf("Service Deletion Error: %s", err))
		}
	}
	white.Println("INFO: ...done")

	white.Println("INFO: Deleting brokered services...")
	for _, svc := range sc.Brokered {
		cmd = exec.Command("cf", "ds", svc.Name, "-f")
		cmd.Stdout = os.Stdout
		err = cmd.Run()
		if err != nil {
			writeErrorAndExit(fmt.Sprintf("Service Deletion Error: %s", err))
		}
	}
	white.Println("INFO: ...done")

	bold.Println("INFO: Destroy succeeded.")
	green.Print("\nSUCCESS\n\n")
	return nil
}

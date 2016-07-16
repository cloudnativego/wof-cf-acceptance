package command

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"gopkg.in/yaml.v2"
)

func execInstalled(file string) bool {
	_, err := exec.LookPath(file)
	if err != nil {
		return false
	}
	return true
}

func pluginInstalled(name string) bool {
	out, err := exec.Command("cf", "plugins").Output()
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(out)))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		if scanner.Text() == name {
			return true
		}
	}
	return false
}

func writeErrorAndExit(message string) {
	bold.Print("ERROR: ")
	white.Println(message)
	red.Print("\nFAILED\n\n")
	os.Exit(1)
}

func readFromFile(source string) (config string, err error) {
	bytes, err := ioutil.ReadFile(source)
	if err != nil {
		return config, err
	}
	config = string(bytes)
	return config, err
}

func parseDeploymentConfig(config string) (dc *DeploymentConfig, err error) {
	dc = &DeploymentConfig{}
	err = yaml.Unmarshal([]byte(config), dc)
	return dc, err
}

func parseSpaceConfiguration(config string) (spaceConfig *SpaceConfiguration, err error) {
	spaceConfig = &SpaceConfiguration{}
	err = yaml.Unmarshal([]byte(config), spaceConfig)
	return spaceConfig, err
}

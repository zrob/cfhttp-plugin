package main

import (
	"fmt"

	"code.cloudfoundry.org/cli/plugin"
	"os/exec"
	"os"
	"strings"
)

type CFHTTPPlugin struct{}

func main() {
	plugin.Start(new(CFHTTPPlugin))
}

func (c *CFHTTPPlugin) Run(cliConnection plugin.CliConnection, args []string) {
	if args[0] == "http" {
		c.http(cliConnection, args[1:])
	}
}

func (c *CFHTTPPlugin) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "cfhttp",
		Version: plugin.VersionType{
			Major: 0,
			Minor: 1,
			Build: 0,
		},
		Commands: []plugin.Command{
			{
				Name:     "http",
				HelpText: "Invokes HTTPie within cf target context",
				UsageDetails: plugin.Usage{
					Usage: "cf http /v3/apps",
				},
			},
		},
	}
}

func (c *CFHTTPPlugin) http(cliConnection plugin.CliConnection, args []string) {
	cmdArgs := []string{"--body", "--ignore-stdin"}
	verb := "get"
	path := args[0]
	httpArgs := args[1:]
	skipSSLValidation, _ := cliConnection.IsSSLDisabled()

	if skipSSLValidation {
		cmdArgs = append(cmdArgs, "--verify=no")
	}

	if isVerb(path) {
		if len(args) < 2 {
			fmt.Println("TODO PRINT USAGE: missing path")
			os.Exit(1)
		}
		verb = path
		path = args[1]
		httpArgs = args[2:]
	}

	endpoint, _ := cliConnection.ApiEndpoint()
	token, _ := cliConnection.AccessToken()

	cmdArgs = append(cmdArgs, verb)
	cmdArgs = append(cmdArgs, fmt.Sprintf("%s%s", endpoint, path))
	cmdArgs = append(cmdArgs, fmt.Sprintf("Authorization:%s", token))
	cmdArgs = append(cmdArgs, httpArgs...)

	cmd := exec.Command("http", cmdArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func isVerb(candidate string) bool {
	return contains([]string{"get", "post", "patch", "put", "head", "delete"}, candidate)
}

func contains(list []string, search string) bool {
	for _, item := range list {
		if strings.EqualFold(item, search) {
			return true
		}
	}
	return false
}
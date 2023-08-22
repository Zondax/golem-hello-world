package main

import (
	"github.com/zondax/golem-hello-world/internal/commands"
	"github.com/zondax/golem-hello-world/internal/conf"
	"github.com/zondax/golem-hello-world/internal/version"
	"strings"
)

import (
	"github.com/zondax/golem/pkg/cli"
)

func main() {
	appName := "golem-hello-world"
	envPrefix := strings.ReplaceAll(appName, "-", "_")

	appSettings := cli.AppSettings{
		Name:        appName,
		Description: "Please override",
		ConfigPath:  "$HOME/.golem-hello-world/",
		EnvPrefix:   envPrefix,
		GitVersion:  version.GitVersion,
		GitRevision: version.GitRevision,
	}

	// Define application level features
	cli := cli.New[conf.Config](appSettings)
	defer cli.Close()

	cli.GetRoot().AddCommand(commands.GetStartCommand(cli))

	cli.Run()
}

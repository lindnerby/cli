package main

import (
	"os"

	"github.com/kyma-project/cli/cmd/kyma"
	"github.com/kyma-project/cli/internal/cli"
)

const lintFail = "test"

func main() {
	command := kyma.NewCmd(cli.NewOptions())

	err := command.Execute()

	if err != nil {
		os.Exit(cli.GetExitCode(err))
	}
}

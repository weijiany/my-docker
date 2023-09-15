package command

import (
	"github.com/urfave/cli"
	"weijiany/docker/src/container"
)

func InitCommand() cli.Command {
	return cli.Command{
		Name:  "init",
		Usage: "Init container process, used to isolate environments and running the user's command",
		Action: func(context *cli.Context) error {
			cmd := context.Args().Get(0)
			return container.RunContainerInitProcess(cmd)
		},
	}
}

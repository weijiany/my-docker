package command

import (
	"fmt"
	"github.com/urfave/cli"
	"weijiany/docker/src/container"
	"weijiany/docker/src/subsystems"
)

func RunCommand() cli.Command {
	input := &subsystems.ResourceConfig{}
	return cli.Command{
		Name:  "run",
		Usage: "Create a container with namespace and cgroup limit. mydocker run -it [command]",
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "it",
				Usage: "enable tty",
			},
			cli.StringFlag{
				Name:        "m",
				Usage:       "set memory",
				Destination: &input.MemoryLimit,
				Value:       "20m",
			},
			cli.StringFlag{
				Name:        "cpu",
				Usage:       "set cpu usage",
				Destination: &input.CpuShare,
				Value:       "100000",
			},
		},
		Action: func(context *cli.Context) error {
			if len(context.Args()) < 1 {
				return fmt.Errorf("missing container command")
			}
			tty := context.Bool("it")
			return container.Run(tty, context.Args(), input)
		},
	}
}

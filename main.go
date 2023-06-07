package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"os"
	"weijiany/docker/src/container"
)

const usage = `is a simple container runtime implementation.
   The purpose of this project is to learn how docker works and how to write a docker by myself
   Enjoy it, just for fun.`

var runCommand = cli.Command{
	Name:  "run",
	Usage: "Create a container with namespace and cgroup limit. mydocker run -it [command]",
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "it",
			Usage: "enable tty",
		},
	},
	Action: func(context *cli.Context) error {
		if len(context.Args()) < 1 {
			return fmt.Errorf("missing container command")
		}
		cmd := context.Args().Get(0)
		tty := context.Bool("it")
		container.Run(tty, cmd)
		return nil
	},
}

func main() {
	app := cli.NewApp()
	app.Name = "mydocker"
	app.Usage = usage

	app.Commands = cli.Commands{
		runCommand,
	}
	app.Before = func(context *cli.Context) error {
		log.SetFormatter(&log.JSONFormatter{})
		log.SetOutput(os.Stdout)
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

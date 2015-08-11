package main

import (
	"log"
	"os"

	"github.com/codegangsta/cli"
)

const (
	Version = "1.0"
)

func main() {
	log.SetFlags(0)
	app := cli.NewApp()
	app.Name = "Benchmark connection"
	app.Version = Version
	app.Author = "hitsumabushi"
	app.Email = "053023.math@gmail.com"
	app.Commands = Commands
	app.EnableBashCompletion = true
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name: "endport, ep",
		},
		cli.IntFlag{
			Name: "beginport, bp",
		},
	}
	app.Run(os.Args)
}

var Commands = []cli.Command{
	commandServer,
	commandClient,
}

var commandServer = cli.Command{
	Name:    "server",
	Aliases: []string{"s"},
	Usage:   "Launch Server",
	Action:  doServer,
	Flags:   []cli.Flag{},
}

var commandClient = cli.Command{
	Name:    "client",
	Aliases: []string{"c"},
	Usage:   "Launch Client",
	Action:  doClient,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name: "host",
		},
		cli.StringFlag{
			Name: "signature, i",
		},
	},
}

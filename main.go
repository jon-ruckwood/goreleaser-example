package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	app := cli.NewApp()
	app.Name = "goreleaser-example"
	app.Usage = "Simple command-line utility"
	app.Version = makeVersion()
	app.Author = "n/a"
	app.Commands = []cli.Command{
		cli.Command{
			Name:        "hello",
			Usage:       "Prints a default message",
			Description: "Prints a default message in response to your greeting.",
			Flags:       nil,
			Action: func(c *cli.Context) {
				println("Hello, world!")
			},
		},
	}

	app.Run(os.Args)
}

func makeVersion() string {
	return fmt.Sprintf("%v (build: %v, on: %v)", version, commit, date)
}

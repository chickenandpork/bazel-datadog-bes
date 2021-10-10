package main

import (
	"fmt"
	"log"
	"os"

	//lib "github.com/chickenandpork/bazel-datadog-bes/lib"
	cli "github.com/urfave/cli/v2"
)

const (
	AuthToken  = "auth_token"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{Name: AuthToken, Aliases: []string{"t"}, Usage: "Auth Token, ie \"\"", EnvVars: []string{"DD_AUTHTOKEN"}},
		},

		Name:  "datadog-bes",
		Usage: "Datadog integration for Bazel Build Event Service",
		Action: func(c *cli.Context) error {
			fmt.Printf("running: AT: [%s]\n", c.String(AuthToken))
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

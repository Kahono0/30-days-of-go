package main

import (
    "log"
    "os"

    cli "github.com/urfave/cli/v2"
)

func main() {
    app := &cli.App{
        Flags: []cli.Flag{
            &cli.StringFlag{
                Name:    "lang",
                Aliases: []string{"l"},
                Value:   "english",
                Usage:   "language for the greeting",
                EnvVars: []string{"LEGACY_COMPAT_LANG", "APP_LANG", "LANG"},
            },
        },
	  Action: func(c *cli.Context) error {
		name := "someone"
		if c.NArg() > 0 {
		    name = c.Args().Get(0)
		}
		if c.String("lang") == "spanish" {
		    log.Printf("Hola %s!", name)
		} else {
		    log.Printf("Hello %s!", name)
		}
		return nil
	  },
    }

    if err := app.Run(os.Args); err != nil {
        log.Fatal(err)
    }
}

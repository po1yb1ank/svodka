package main

import (
	"github.com/po1yb1ank/svodka/modules"
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name:  "currency",
		Usage: "prints today's currency",
		Action: func(c *cli.Context) error {
			modules.CurrencyStats()
			modules.WeatherStats()
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

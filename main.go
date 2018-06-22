package main

import (
	"log"
	"os"

	"github.com/asicsdigital/dudewheresmy/commands/hostip"
	"gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Name = "dudewheresmy"
	app.Usage = "find things you're looking for"

	app.Commands = []cli.Command{
		hostip.Command(),
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

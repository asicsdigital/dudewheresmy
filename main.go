package main

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/asicsdigital/dudewheresmy/commands/hostip"
	"gopkg.in/urfave/cli.v1"
)

const (
	appname    string = "dudewheresmy"
	appversion string = "v0.1.0"
	appusage   string = "find things you're looking for"
)

func main() {
	log.SetLevel(log.WarnLevel)

	app := cli.NewApp()
	app.Name = appname
	app.Version = appversion
	app.Usage = appusage

	app.Commands = []cli.Command{
		hostip.Command(),
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

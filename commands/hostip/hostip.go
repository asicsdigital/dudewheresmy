package hostip

import (
	"fmt"
	"net"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"
	"gopkg.in/urfave/cli.v1"
)

const (
	dockerlocal  string = "host.docker.internal"
	localhost    string = "127.0.0.1"
	metadatapath string = "local-ipv4"
	timeoutsec   = 2
)

func Command() cli.Command {
	c := cli.Command{
		Name:      "hostip",
		ShortName: "i",
		Aliases:   []string{"ip", "host"},
		Usage:     "find IP address of process host",
		Action:    action,
	}

	return c
}

func action(*cli.Context) error {
	log.SetLevel(log.WarnLevel)

	// the channel only needs to be big enough for one response
	hostips := make(chan string, 1)
	hostip := localhost

	go fromEc2Metadata(hostips)
	go fromDockerInternal(hostips)

	select {
	case res := <-hostips:
		hostip = res
	case <-time.After(timeoutsec * time.Second):
		log.Printf("timeout after %d seconds", timeoutsec)
	}

	err := parseAndPrint(hostip)

	return err
}

func parseAndPrint(i string) error {
	parsed := net.ParseIP(i)

	var err error

	if parsed != nil {
		_, err = fmt.Printf("%v\n", parsed)
	} else {
		err = fmt.Errorf("unable to parse %v as an IP address", i)
	}

	return err
}

func fromEc2Metadata(c chan string) {
	sess := session.Must(session.NewSession())

	client := ec2metadata.New(sess)

	if client.Available() {
		hostip, err := client.GetMetadata(metadatapath)

		if err != nil {
			log.Panic(err)
		} else {
			c <- hostip
		}
	}
}

func fromDockerInternal(c chan string) {
	hostips, err := net.LookupHost(dockerlocal)

	if err == nil {
		c <- hostips[0]
	}
}

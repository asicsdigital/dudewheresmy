package hostip

import (
	"fmt"
	"log"
	"net"

	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"
	"gopkg.in/urfave/cli.v1"
)

const (
	dockerlocal  string = "host.docker.internal"
	localhost    string = "127.0.0.1"
	metadatapath string = "local-ipv4"
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
	hostip, err := fromEc2Metadata()

	if err != nil {
		log.Print(err)
	}

	hostip, err = fromDockerInternal()

	if err != nil {
		log.Print(err)
	}

	if hostip == "" {
		hostip = localhost
	}

	parsed := net.ParseIP(hostip)

	if parsed != nil {
		_, err = fmt.Printf("%v\n", parsed)
	} else {
		err = fmt.Errorf("unable to parse %v as an IP address", hostip)
	}

	return err
}

func fromEc2Metadata() (string, error) {
	sess := session.Must(session.NewSession())

	client := ec2metadata.New(sess)

	if client.Available() {
		hostip, err := client.GetMetadata(metadatapath)

		return hostip, err
	} else {
		return "", fmt.Errorf("EC2 metadata service not available")
	}
}

func fromDockerInternal() (string, error) {
	hostips, err := net.LookupHost(dockerlocal)

	if err != nil {
		return "", err
	} else {
		return hostips[0], err
	}
}

package cmd

import (
	"fmt"
	"net"
	"time"

	"github.com/spf13/cobra"
	jww "github.com/spf13/jwalterweatherman"

	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"
)

const (
	dockerlocal  string = "host.docker.internal"
	localhost    string = "127.0.0.1"
	metadatapath string = "local-ipv4"
	timeoutsec          = 2
)

func init() {
	rootCmd.AddCommand(hostipCmd)
}

var hostipCmd = &cobra.Command{
	Use:   "hostip",
	Short: "find IP address of process host",
	Run: func(cmd *cobra.Command, args []string) {
		hostIp()
	},
}

func hostIp() {

	// the channel only needs to be big enough for one response
	hostips := make(chan string, 1)
	hostip := localhost

	go fromEc2Metadata(hostips)
	go fromDockerInternal(hostips)

	select {
	case res := <-hostips:
		hostip = res
	case <-time.After(timeoutsec * time.Second):
		jww.DEBUG.Printf("timeout after %d seconds", timeoutsec)
	}

	err := parseAndPrint(hostip)

	if err != nil {
		jww.FATAL.Println(err)
	}
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
			jww.FATAL.Println(err)
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

package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Generate returns the CLI app ready to be executed
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "CLI Application"
	app.Usage = "Get server IPs and Hostnames from the Internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name: "host",
			Value: "google.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name: "ip",
			Usage: "Get IPs from addreses from the Internet",
			Flags: flags,
			Action: getIps,
		},
		{
			Name: "hostnames",
			Usage: "Get hostnames from the Internet",
			Flags: flags,
			Action: getHostnames,
		},
	}

	return app
}

func getIps(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func getHostnames(c *cli.Context) {
	host := c.String("host")

	hostnames, err := net.LookupNS(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, hostname := range hostnames {
		fmt.Println(hostname.Host)
	}
}
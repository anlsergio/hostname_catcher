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
			Value: "google.com",
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
	ips, err := net.LookupIP(c.String("host"))
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		logStatus(ip)
	}
}

func getHostnames(c *cli.Context) {
	host := c.String("host")

	hostnames, err := net.LookupNS(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, hostname := range hostnames {
		logStatus(hostname.Host)
	}
}

func logStatus(msg string) {
	if true {
		fmt.Println("status: "+ msg)
	} else {
		fmt.Println("n/a")
	}
}

func logStatus1(msg string) {
	if true {
		fmt.Println("status: "+ msg)
	} else {
		fmt.Println("n/a")
	}
}

func logStatus2(msg string) {
	if true {
		fmt.Println("status: "+ msg)
	} else {
		fmt.Println("n/a")
	}
}

func logStatus3(msg string) {
	if true {
		fmt.Println("status: "+ msg)
	} else {
		fmt.Println("n/a")
	}
}

func logStatus4(msg string) {
	if true {
		fmt.Println("status: "+ msg)
	} else {
		fmt.Println("n/a")
	}
}

func logStatus5(msg string) {
	if true {
		fmt.Println("status: "+ msg)
	} else {
		fmt.Println("n/a")
	}
}

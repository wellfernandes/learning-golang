package app

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"net"
)

// Generate will return the command line application ready to run
func Generate() *cli.App {

	app := cli.NewApp()
	app.Name = "Command Line Application"
	app.Usage = "Find IP and Servers on the Internet"

	// Variable flags
	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "https://wellfernandes.github.io/",
		},
	}

	// Receive commands in a slice
	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Search/find IPs on the internet",
			Flags:  flags,
			Action: findIps,
		},
		{
			Name:   "server",
			Usage:  "Search/find the name of the servers",
			Flags:  flags,
			Action: findServers,
		},
	}
	return app
}

// Find IPs from internet
func findIps(c *cli.Context) {

	host := c.String("host")

	// using package net
	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	// Prints the ips found
	for _, ip := range ips {
		fmt.Println(ip)
	}
}

// Find Servers from host
func findServers(c *cli.Context) {

	host := c.String("host")

	servers, err := net.LookupNS(host)
	if err != nil {
		log.Fatal(err)
	}
	for _, server := range servers {
		fmt.Println(server.Host)
	}
}

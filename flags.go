package main

import "gopkg.in/urfave/cli.v2"

func setFlags() []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:    "verbose",
			Value:   false,
			Aliases: []string{"v"},
			Usage:   "Verbose output",
		},
		&cli.BoolFlag{
			Name:    "quiet",
			Value:   false,
			Aliases: []string{"q"},
			Usage:   "Quiet",
		},
		&cli.StringFlag{
			Name:    "machine",
			Value:   "",
			Aliases: []string{"m"},
			Usage:   "Execute on machine",
		},
		&cli.BoolFlag{
			Name:    "all",
			Value:   false,
			Aliases: []string{"a"},
			Usage:   "Execute on all machines",
		},
		&cli.StringFlag{
			Name:    "group",
			Value:   "",
			Aliases: []string{"g"},
			Usage:   "Execute on group member",
		},
		&cli.StringFlag{
			Name:    "file",
			Value:   "",
			Aliases: []string{"f"},
			Usage:   "Use the file as a list of machines",
		},
		&cli.StringFlag{
			Name:    "remoteshell",
			Value:   "",
			Aliases: []string{"r"},
			Usage:   "Execute using shell (rsh/ssh)",
		},
		&cli.StringFlag{
			Name:    "remoteshellopt",
			Value:   "",
			Aliases: []string{"o"},
			Usage:   "Option to give to shell",
		},
		&cli.BoolFlag{
			Name:    "wait-shell",
			Value:   false,
			Aliases: []string{"w"},
			Usage:   "Sequentially execute shell",
		},
		&cli.BoolFlag{
			Name:    "concurrent-shell",
			Value:   false,
			Aliases: []string{"c"},
			Usage:   "Execute shell concurrently",
		},
		&cli.BoolFlag{
			Name:    "show-machine-names",
			Value:   false,
			Aliases: []string{"M"},
			Usage:   "Prepend the hostname on output",
		},
		&cli.BoolFlag{
			Name:    "hide-machine-names",
			Value:   false,
			Aliases: []string{"H"},
			Usage:   "Do not prepend host name on output",
		},
	}
}

package main

import (
	"container/list"
	"fmt"
	"path/filepath"

	"github.com/grubernaut/gdsh/dsh"

	"gopkg.in/urfave/cli.v2"
)

// ConfDir is the default configuration directory
var ConfDir = "/usr/local/etc/dsh"

// ConfPath is the filepath to the default machine list file
var ConfPath = filepath.Join(ConfDir, "machines.list")

func run(c *cli.Context) error {
	// TODO: Not actually an rule. Need to remove
	if c.NArg() < 1 {
		fmt.Printf("Error: expected an execution argument\n")
		return cli.ShowAppHelp(c)
	}
	// Create ExecOpts
	opts := dsh.ExecOpts{}
	// Create an empty machineList, and initialize struct
	machineList := list.New()
	opts.MachineList = machineList

	// Set opts.Verbose output
	opts.Verbose = false
	if c.Bool("verbose") {
		fmt.Printf("Verbose flag on\n")
		opts.Verbose = true
	}
	if c.Bool("quiet") {
		fmt.Printf("verbose flag off\n")
	}

	// Build up exec opts
	if c.Bool("all") {
		if opts.Verbose {
			fmt.Printf("Adding all machines to the list\n")
		}
		if err := opts.ReadMachineList(ConfPath); err != nil {
			_ = cli.Exit(fmt.Sprintf("Error adding all machines: %s", err), 1)
		}
	}
	if c.String("group") != "" {
		if opts.Verbose {
			fmt.Printf("Adding group %s to the list\n", c.String("group"))
		}
	}
	if c.String("file") != "" {
		if opts.Verbose {
			fmt.Printf("Adding file %s to the list\n", c.String("file"))
		}
		if err := opts.ReadMachineList(c.String("file")); err != nil {
			_ = cli.Exit(fmt.Sprintf("Error adding machines from %s: %s",
				c.String("file"), err), 1)
		}
	}
	if c.Bool("show-machine-names") {
		if opts.Verbose {
			fmt.Printf("Show machine names on output\n")
		}
		opts.ShowNames = true
	}
	if c.Bool("hide-machine-names") {
		if opts.Verbose {
			fmt.Printf("Don't show machine names on output\n")
		}
		opts.ShowNames = false
	}
	if c.String("machine") != "" {
		if opts.Verbose {
			fmt.Printf("Adding machine %s to list\n", c.String("machine"))
		}
	}
	if c.String("remoteshell") != "" {
		if opts.Verbose {
			fmt.Printf("Using %s as remote shell\n", c.String("remoteshell"))
		}
		if c.String("remoteshell") == "ssh" || c.String("remoteshell") == "rsh" {
			opts.RemoteShell = c.String("remoteshell")
		} else {
			return cli.Exit("gdsh only supports 'ssh' and 'rsh' currently", 1)
		}
	}
	if c.String("remoteshellopt") != "" {
		if opts.Verbose {
			fmt.Printf("Adding [%s] to shell options\n", c.String("remoteshellopt"))
		}
		opts.RemoteCommandOpts = c.String("remoteshellopt")
	}
	if c.Bool("wait-shell") {
		if opts.Verbose {
			fmt.Printf("Wait for shell to finish executing\n")
		}
		opts.ConcurrentShell = false
	}
	if c.Bool("concurrent-shell") {
		if opts.Verbose {
			fmt.Printf("Do not wait for shell to finish\n")
		}
		opts.ConcurrentShell = true
	}
	// Set remote command to execute
	for _, v := range c.Args().Slice() {
		if opts.RemoteCommand == "" && v != "" {
			opts.RemoteCommand = v
			continue
		}
		opts.RemoteCommand = fmt.Sprintf("%s %s", opts.RemoteCommand, v)
	}

	// Execute
	if err := opts.Execute(); err != nil {
		return cli.Exit(fmt.Sprintf("Error executing: %s", err), 1)
	}
	return nil
}

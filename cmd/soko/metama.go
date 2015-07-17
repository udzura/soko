package main

import (
	"flag"
	"os"

	"github.com/udzura/soko"
)

var serverID string

func main() {
	const usage = "Target server's ID to get/put/delete. Defaults to cloud-init's server ID"

	// FIXME: get default server id via cloud directory
	flag.StringVar(&serverID, "server-id", soko.CloudServerID(), usage)

	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	command := args[0]
	realArgs := args[1:]

	runner := &soko.Runner{
		ServerID: serverID,
	}
	runner.Run(command, realArgs)
}

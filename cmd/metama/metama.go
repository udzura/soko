package main

import (
	"flag"

	"github.com/udzura/metama"
)

var serverID string

func main() {
	const usage = "Target server's ID to get/put/delete. Defaults to cloud-init's server ID"

	// FIXME: get default server id via cloud directory
	// flag.StringVar(&string, "server-id", metama.CloudServerID())
	flag.StringVar(&serverID, "server-id", "dummy", usage)

	flag.Parse()

	args := flag.Args()
	command := args[0]
	realArgs := args[1:]

	runner := &metama.Runner{
		ServerID: serverID,
	}
	runner.Run(command, realArgs)
}

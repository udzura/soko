package metama

import (
	"flag"
	"fmt"
	"os"
)

type Runner struct {
	ServerID string

	backend Backend
}

func (r *Runner) Run(subcommand string, args []string) {
	r.backend = FindBackend("consul://dummy:8500")

	switch subcommand {
	case "get":
		checkArgSizeOf(args, 1)
		key := args[0]
		r.Get(key)
	case "put":
		checkArgSizeOf(args, 2)
		key := args[0]
		value := args[1]
		r.Put(key, value)
	case "delete":
		checkArgSizeOf(args, 1)
		key := args[0]
		r.Delete(key)

	case "join":
		fmt.Fprintf(os.Stderr, "join is not yet implemented...\n")
		flag.Usage()

	case "version":
		fmt.Printf("version v%s\n", Version)

	default:
		flag.Usage()
	}
}

func (r *Runner) Get(key string) {
	ret, err := r.backend.Get(r.ServerID, key)
	if err != nil {
		panic(err)
	}
	fmt.Println(ret)
}

func (r *Runner) Put(key string, value string) {
	err := r.backend.Put(r.ServerID, key, value)
	if err != nil {
		panic(err)
	}
	fmt.Println("OK")
}

func (r *Runner) Delete(key string) {
	err := r.backend.Delete(r.ServerID, key)
	if err != nil {
		panic(err)
	}
	fmt.Println("OK")
}

func checkArgSizeOf(args []string, size int) {
	if len(args) != size {
		fmt.Fprintf(os.Stderr, "Argument size is mismatch for subcommand: %v\n", flag.Args())
		flag.Usage()
		os.Exit(1)
	}
}

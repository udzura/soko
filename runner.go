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
		key := args[0]
		r.Get(key)
	case "put":
		key := args[0]
		value := args[1]
		r.Put(key, value)
	case "delete":
		key := args[0]
		r.Delete(key)

	case "join":
		fmt.Fprintf(os.Stderr, "join is not yet implemented...\n")
		flag.Usage()

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

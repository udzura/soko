package soko

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
	switch subcommand {
	case "open", "config":
		config, err := NewConfig(args[0], args[1:])
		if err != nil {
			panic(err)
		}

		backend, err := FindBackend(config)
		if err != nil {
			panic(err)
		}
		r.backend = backend
		r.Save()
	case "get":
		r.prepare()
		checkArgSizeOf(args, 1)
		key := args[0]
		r.Get(key)
	case "put":
		r.prepare()
		checkArgSizeOf(args, 2)
		key := args[0]
		value := args[1]
		r.Put(key, value)
	case "delete":
		r.prepare()
		checkArgSizeOf(args, 1)
		key := args[0]
		r.Delete(key)

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

func (r *Runner) Save() {
	err := r.backend.Save()
	if err != nil {
		panic(err)
	}
	fmt.Printf("OK: Write config to %s\n", defaultConfigPath)
}

func (r *Runner) prepare() {
	config, err := DefaultConfig()
	if err != nil {
		panic(err)
	}

	backend, err := FindBackend(config)
	if err != nil {
		panic(err)
	}
	r.backend = backend
}

func checkArgSizeOf(args []string, size int) {
	if len(args) != size {
		fmt.Fprintf(os.Stderr, "Argument size is mismatch for subcommand: %v\n", flag.Args())
		flag.Usage()
		os.Exit(1)
	}
}

package soko

import (
	"flag"
	"fmt"
)

func init() {
	orig := flag.Usage
	flag.Usage = func() {
		orig()
		fmt.Printf(`Valid subcommands:
open   - initialize backend config
get    - get metadata value
put    - put metadata value with key
delete - delete metadata value
`)
	}
}

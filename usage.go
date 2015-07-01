package metama

import (
	"flag"
	"fmt"
)

func init() {
	orig := flag.Usage
	flag.Usage = func() {
		orig()
		fmt.Printf(`Valid subcommands:
join   - initialize backend config
get    - get metadata value
put    - put metadata value with key
delete - delete metadata value
`)
	}
}

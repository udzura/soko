package soko

import (
	"fmt"
	"os"
)

func sayEmpty(key string) {
	fmt.Fprintf(os.Stderr, "Value for %s seems to be empty.\n", key)
}

func toStringMap(m map[string]interface{}) map[string]string {
	newMap := make(map[string]string, 0)
	for k, v := range m {
		if newV, ok := v.(string); ok {
			newMap[k] = newV
		}
	}

	return newMap
}

package metama

import (
	"io/ioutil"
	"os"
	"strings"
)

const cloudInstanceIDFile = "/var/lib/cloud/data/instance-id"

func CloudServerID() string {
	if _, err := os.Stat(cloudInstanceIDFile); os.IsNotExist(err) {
		return ""
	} else {
		f, err := os.Open(cloudInstanceIDFile)
		if err != nil {
			panic(err)
		}
		d, err := ioutil.ReadAll(f)
		if err != nil {
			panic(err)
		}

		return strings.TrimSpace(string(d))
	}
}

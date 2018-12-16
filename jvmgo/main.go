package main

import (
	"os"

	"github.com/INFIA-FOUNDATION/TEST_VM/cmdline"
)

func main() {
	cmd, err := cmdline.ParseCommand(os.Args)
	if err != nil {
		cmdline.PrintUsage()
	} else {
		startJVM(cmd)
	}
}

package main

import (
	"fmt"
	"os"

	"github.com/micnncim/octo/pkg/cmd"
)

func main() {
	if err := cmd.NewCommand().Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

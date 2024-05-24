package main

import (
	"os"

	"github.com/vinicius507/memoscli/cmd"
)

func main() {
	cmd := cmd.NewRootCmd()
	err := cmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

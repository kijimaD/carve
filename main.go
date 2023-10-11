package main

import (
	"os"

	"github.com/kijimaD/carve/cmd"
)

func main() {
	cmd := cmd.New(os.Stdout)
	err := cmd.Execute(os.Args)
	if err != nil {
		panic(err)
	}
}

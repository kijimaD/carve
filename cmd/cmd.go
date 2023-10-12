package cmd

import (
	"errors"
	"flag"
	"io"
	"path/filepath"

	carve "github.com/kijimaD/carve/pkg"
)

type CLI struct {
	Out io.Writer
}

var (
	NotEnoughArgumentError = errors.New("Not enough arguments. Expect greater then 3 arguments.")
)

func New(out io.Writer) *CLI {
	return &CLI{
		Out: out,
	}
}

func (cli *CLI) Execute(args []string) error {
	flag.Parse()

	if len(args) <= 3 {
		return NotEnoughArgumentError
	}

	gitpath := args[1]
	oldversion := args[2]
	files := args[3:]
	newversion, err := carve.GetNewTag(filepath.Join(gitpath, ".git"))
	if err != nil {
		return err
	}

	carve.Replacewalk(files, oldversion, newversion)

	return nil
}

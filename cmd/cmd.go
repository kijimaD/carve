package cmd

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"path/filepath"

	carve "github.com/kijimaD/carve/pkg"
)

type CLI struct {
	Out io.Writer
}

var (
	NotEnoughArgumentError = errors.New("Not enough arguments. Expect greater then 2 arguments.")
)

func New(out io.Writer) *CLI {
	return &CLI{
		Out: out,
	}
}

func (cli *CLI) Execute(args []string) error {
	flag.Parse()

	if len(args) <= 2 {
		return NotEnoughArgumentError
	}

	gitpath := args[1]
	files := args[2:]
	newtag, err := carve.GetLatestTag(filepath.Join(gitpath, ".git"))
	if err != nil {
		return err
	}
	oldtag, err := carve.GetOldTag()
	// タグファイルがない場合は、最新タグでタグファイルを作成する
	// TODO: ファイルが存在しないときは、という感じにする
	if oldtag == "" {
		fmt.Fprintf(cli.Out, "file `%s` is not found, created...\n", carve.Versionfile)
		carve.PutTagFile(".")
		oldtag, err = carve.GetLatestTag(filepath.Join(gitpath, ".git"))
		if err != nil {
			return err
		}
	}

	carve.Replacewalk(files, oldtag, newtag)
	fmt.Fprintf(
		cli.Out,
		"%s -> %s\n",
		oldtag,
		newtag,
	)
	carve.PutTagFile(".")

	return nil
}

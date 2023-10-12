package carve

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

func GetNewTag(repopath string) (string, error) {
	r, err := git.PlainOpen(repopath)
	if err != nil {
		log.Fatal(err)
	}

	tagIter, err := r.Tags()
	if err != nil {
		return "", err
	}

	var version string
	// MEMO: 古い順にイテレートされ、ループの最後で最新のバージョンが入る
	tagIter.ForEach(func(ref *plumbing.Reference) error {
		version = ref.Name().Short()
		return nil
	})
	return version, nil
}

func Replacewalk(targetpath []string, old string, new string) error {
	rootDir := "."
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		for _, tp := range targetpath {
			if !info.IsDir() && path == tp {
				if err := replacefile(path, old, new); err != nil {
					return err
				}
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func replacefile(filepath string, old string, new string) error {
	f, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	replaced := strings.ReplaceAll(string(b), old, new)
	err = ioutil.WriteFile(filepath, []byte(replaced), os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

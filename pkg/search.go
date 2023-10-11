package carve

import (
	"log"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

func getVersion(repopath string) (string, error) {
	r, err := git.PlainOpen(repopath)
	if err != nil {
		log.Fatal(err)
	}

	tagIter, err := r.Tags()
	if err != nil {
		return "", err
	}

	var version string
	tagIter.ForEach(func(ref *plumbing.Reference) error {
		version = ref.Name().Short()
		return nil
	})
	return version, nil
}

package carve

import (
	"fmt"
	"log"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

func getVersion() (string, error) {
	r, err := git.PlainOpen("../.git")
	if err != nil {
		log.Fatal(err)
	}

	// タグを一覧表示
	tagIter, err := r.Tags()
	if err != nil {
		log.Fatal(err)
	}

	tagIter.ForEach(func(ref *plumbing.Reference) error {
		tagHash := ref.Hash()
		fmt.Printf("タグ名: %s\n", ref.Name().Short())
		fmt.Printf("タグのコミットハッシュ: %s\n", tagHash.String())

		// コミットオブジェクトを取得
		commitObj, err := r.CommitObject(tagHash)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("コミットメッセージ: %s\n", commitObj.Message)

		return nil
	})

	commitHash := plumbing.NewHash("your_commit_hash")
	commitObj, err := r.CommitObject(commitHash)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("特定のコミットメッセージ: %s\n", commitObj.Message)

	return "", nil
}

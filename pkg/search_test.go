package carve

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/stretchr/testify/assert"
)

// gitリポジトリを作成して確かめる
func TestVersion(t *testing.T) {
	err := os.RemoveAll("./.git")
	if err != nil {
		log.Fatal(err)
	}

	repoPath := "."

	// 指定したディレクトリに新しいGitリポジトリを初期化
	r, err := git.PlainInit(repoPath, false)
	if err != nil {
		log.Fatal(err)
	}

	// ワーキングツリーからステージングエリアにファイルを追加（必要な場合）
	w, err := r.Worktree()
	if err != nil {
		log.Fatal(err)
	}

	// ステージングエリアにファイルを追加
	_, err = w.Add("dummy")
	if err != nil {
		log.Fatal(err)
	}

	// 新しいコミットを作成
	_, err = w.Commit("dummy commit", &git.CommitOptions{
		Author: &object.Signature{
			Name:  "Your Name",
			Email: "your.email@example.com",
			When:  time.Now(),
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	// 最新コミット
	h, err := r.Head()
	// タグを追加
	_, err = r.CreateTag("v1.0.0", h.Hash(), &git.CreateTagOptions{
		Tagger: &object.Signature{
			Name:  "Your Name",
			Email: "your.email@example.com",
			When:  time.Now(),
		},
		Message: "tag message",
	})
	if err != nil {
		log.Fatal(err)
	}

	_, err = w.Commit("dummy commit2", &git.CommitOptions{
		Author: &object.Signature{
			Name:  "Your Name",
			Email: "your.email@example.com",
			When:  time.Now(),
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	h, err = r.Head()
	// タグを追加
	_, err = r.CreateTag("v2.0.0", h.Hash(), &git.CreateTagOptions{
		Tagger: &object.Signature{
			Name:  "Your Name",
			Email: "your.email@example.com",
			When:  time.Now(),
		},
		Message: "tag message",
	})
	if err != nil {
		log.Fatal(err)
	}

	v, err := GetNewTag("./.git")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, "v2.0.0", v)

	err = os.RemoveAll("./.git")
	if err != nil {
		log.Fatal(err)
	}
}

func TestReplacewalk(t *testing.T) {
	Replacewalk([]string{"dummy1", "dummy2"}, "xxxx", "yyyy")
}

func TestReplacefile(t *testing.T) {
	replacefile("dummy", "xxxx", "yyyy")
}

func TestPlaceTag(t *testing.T) {
	err := PlaceTag()
	if err != nil {
		t.Error(err)
	}
}

func TestGetOldTag(t *testing.T) {
	oldtag, err := GetOldTag()
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, "v2.0.0", oldtag)
}

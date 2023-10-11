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
	tagName := "v1.0.0"             // タグ名
	tagMessage := "Initial release" // タグメッセージ

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
	_, err = r.CreateTag(tagName, h.Hash(), &git.CreateTagOptions{
		Tagger: &object.Signature{
			Name:  "Your Name",
			Email: "your.email@example.com",
			When:  time.Now(),
		},
		Message: tagMessage,
	})
	if err != nil {
		log.Fatal(err)
	}

	v, err := getVersion("./.git")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, "v1.0.0", v)

	err = os.RemoveAll("./.git")
	if err != nil {
		log.Fatal(err)
	}
}

func TestSearch(t *testing.T) {
	search([]string{"dummy1", "dummy2"}, "xxxx", "yyyy")
}

func TestReplacefile(t *testing.T) {
	replacefile("dummy", "xxxx", "yyyy")
}

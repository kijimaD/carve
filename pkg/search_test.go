package carve

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/stretchr/testify/assert"
)

func makeGitRepo(t *testing.T) {
	t.Helper()
	err := os.RemoveAll("./.git")
	if err != nil {
		log.Fatal(err)
	}

	tempfile, err := ioutil.TempFile(".", "")
	if err != nil {
		t.Error(err)
	}
	defer os.RemoveAll(tempfile.Name())

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
	_, err = w.Add(tempfile.Name())
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
		t.Error(err)
	}
}

func TestGetNewTag(t *testing.T) {
	makeGitRepo(t)
	defer os.RemoveAll("./.git")

	v, err := GetNewTag("./.git")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, "v2.0.0", v)
}

func TestReplacewalk(t *testing.T) {
	tempfile, err := ioutil.TempFile(".", "")
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(tempfile.Name())
	ioutil.WriteFile(tempfile.Name(), []byte("xxxx yyyy zzzz"), os.ModePerm)
	Replacewalk([]string{tempfile.Name()}, "x", "y")

	b, err := ioutil.ReadAll(tempfile)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, "yyyy yyyy zzzz", string(b))
}

func TestReplacefile(t *testing.T) {
	tempfile, err := ioutil.TempFile(".", "")
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(tempfile.Name())
	ioutil.WriteFile(tempfile.Name(), []byte("xxxx yyyy zzzz"), os.ModePerm)
	replacefile(tempfile.Name(), "x", "y")

	b, err := ioutil.ReadAll(tempfile)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, "yyyy yyyy zzzz", string(b))
}

func TestPutTagFile(t *testing.T) {
	makeGitRepo(t)
	defer os.RemoveAll("./.git")
	defer os.RemoveAll(filepath.Join("./", Versionfile))

	err := PutTagFile(".")
	if err != nil {
		t.Error(err)
	}

	data, err := ioutil.ReadFile(filepath.Join("./", Versionfile))
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, "v2.0.0", string(data))
}

func TestGetOldTag(t *testing.T) {
	makeGitRepo(t)
	defer os.RemoveAll("./.git")
	defer os.RemoveAll(filepath.Join("./", Versionfile))

	err := PutTagFile(".")
	if err != nil {
		t.Error(err)
	}

	oldtag, err := GetOldTag()
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, "v2.0.0", oldtag)
}

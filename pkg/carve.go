package carve

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

type carve struct {
	RepoPath string // TODO: 取り回しやすいようにpathではなくRepoを渡す
	OldTag   string
	NewTag   string
}

const Versionfile = ".versions"

func GetLatestTag(repopath string) (string, error) {
	r, err := git.PlainOpen(repopath)
	if err != nil {
		return "", err
	}
	tags, err := getTags(r)
	if err != nil {
		return "", err
	}
	sortTagsByDate(tags, r)
	// 日付でソートして最新のタグを返す
	return tags[0].Name().Short(), nil
}

func getTags(repo *git.Repository) ([]*plumbing.Reference, error) {
	tags, err := repo.Tags()
	if err != nil {
		return nil, err
	}
	var tagList []*plumbing.Reference
	err = tags.ForEach(func(tag *plumbing.Reference) error {
		tagList = append(tagList, tag)
		return nil
	})
	return tagList, err
}

func getCommitTime(tag *plumbing.Reference, repo *git.Repository) (time.Time, error) {
	hash := tag.Hash()
	tagObj, err := repo.TagObject(hash)
	if err != nil {
		return time.Time{}, err
	}

	return tagObj.Tagger.When, nil
}

func sortTagsByDate(tags []*plumbing.Reference, repo *git.Repository) error {
	var e error
	sort.Slice(tags, func(i, j int) bool {
		timeI, err := getCommitTime(tags[i], repo)
		if err != nil {
			e = err
		}
		timeJ, err := getCommitTime(tags[j], repo)
		if err != nil {
			e = err
		}
		return timeI.After(timeJ)
	})
	return e
}

func Replacewalk(targetpaths []string, old string, new string) error {
	// MEMO: パスを正規化する。`./`を含まない形に統一しないとマッチしない
	cleanpaths := make([]string, len(targetpaths))
	for i, v := range targetpaths {
		cleanpaths[i] = canonicalPath(v)
	}

	const rootDir = "."
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		for _, cp := range cleanpaths {
			if !info.IsDir() && path == cp {
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

// .versionsを配置する
func PutTagFile(basepath string) error {
	tag, err := GetLatestTag(".")
	if err != nil {
		return err
	}

	fp, err := os.Create(filepath.Join(basepath, Versionfile))
	if err != nil {
		return err
	}
	defer fp.Close()

	fp.WriteString(tag)
	return nil
}

// .versionsからタグを取得する
func GetOldTag() (string, error) {
	data, err := ioutil.ReadFile(Versionfile)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

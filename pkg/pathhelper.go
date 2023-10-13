package carve

import (
	"os"
	"path/filepath"
	"strings"
)

// パスを正規化する
func canonicalPath(path string) string {
	splitPath := strings.Split(path, string(os.PathSeparator))
	// 絶対パスの場合は先頭文字が空文字になるので追加する
	if filepath.IsAbs(path) {
		splitPath[0] = "/"
	}
	return filepath.Join(splitPath...)
}

package utils

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

/**
统一文件的后缀 比如 go实战-mPandaer.md java实战-mPandaer.md => go实战.md java实战.md
*/

func RemoveSuffixFromFile(dir string, suffix string) error {
	err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		oldName := info.Name()
		ext := filepath.Ext(oldName)
		newSuffix := suffix + ext
		if !info.IsDir() && strings.HasSuffix(info.Name(), newSuffix) {
			newName := strings.TrimSuffix(oldName, newSuffix) + ext
			newPath := filepath.Join(filepath.Dir(path), newName)
			err := os.Rename(path, newPath)
			if err != nil {
				fmt.Println("文件重命名失败:", info.Name())
				return err
			}
		}
		return nil
	})

	return err
}

package analyzer

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"

	"github.com/Morizz00/codebase-analyzer/utils"
)

type FileStats struct {
	Path     string
	Language string
	Lines    int
	Size     int64
}

func Scan(root string) []FileStats {
	var res []FileStats
	filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return nil
		}
		if d.IsDir() {
			return nil
		}

		if strings.HasPrefix(d.Name(), ".") {
			return nil
		}

		ext := filepath.Ext(path)
		lang := utils.GetLangFromExt(ext)
		if lang == "" {
			return nil
		}
		file, err := os.Open(path)
		if err != nil {
			return nil
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		lines := 0
		for scanner.Scan() {
			lines++
		}
		info, err := file.Stat()
		if err != nil {
			return nil
		}
		size := info.Size()

		res = append(res, FileStats{
			Path:     path,
			Language: lang,
			Lines:    lines,
			Size:     size,
		})
		return nil
	})
	return res

}

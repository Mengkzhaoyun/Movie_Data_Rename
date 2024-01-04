package libs

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func Movie_list(root string) []string {

	result := make([]string, 0)

	if !IsDir(root) {
		log.Printf("Source Dir is not exist.: %s", root)
	}

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if IsVedio(path) && !info.IsDir() {
			result = append(result, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	return result
}

func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func IsVedio(path string) bool {
	if strings.HasSuffix(path, ".mp4") {
		return true
	}
	if strings.HasSuffix(path, ".avi") {
		return true
	}
	if strings.HasSuffix(path, ".mkv") {
		return true
	}
	return false
}

func IsNumber(path string) bool {
	if strings.HasSuffix(strings.ToLower(path), "-l") {
		return true
	}
	if strings.HasSuffix(strings.ToLower(path), "-u") {
		return true
	}
	if strings.HasSuffix(strings.ToLower(path), "-c") {
		return true
	}
	if strings.HasSuffix(strings.ToLower(path), "-lc") {
		return true
	}
	if strings.HasSuffix(strings.ToLower(path), "-uc") {
		return true
	}
	return false
}

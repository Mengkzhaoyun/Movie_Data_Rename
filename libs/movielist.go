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

	err := filepath.Walk(root, func(sPath string, info os.FileInfo, err error) error {
		if IsVedio(sPath) && !info.IsDir() {
			result = append(result, sPath)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	return result
}

func IsDir(sPath string) bool {
	s, err := os.Stat(sPath)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func IsVedio(sPath string) bool {
	if filepath.Ext(sPath) == ".mp4" {
		return true
	}
	if filepath.Ext(sPath) == ".avi" {
		return true
	}
	if filepath.Ext(sPath) == ".mkv" {
		return true
	}
	return false
}

func IsNumber(sPath string) bool {
	if strings.HasSuffix(strings.ToLower(sPath), "-l") {
		return true
	}
	if strings.HasSuffix(strings.ToLower(sPath), "-u") {
		return true
	}
	if strings.HasSuffix(strings.ToLower(sPath), "-c") {
		return true
	}
	if strings.HasSuffix(strings.ToLower(sPath), "-lc") {
		return true
	}
	if strings.HasSuffix(strings.ToLower(sPath), "-uc") {
		return true
	}
	return false
}

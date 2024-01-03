package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mengkzhaoyun/movie_data_rename/libs"
)

var Source = "./"

func main() {
	Source, _ = os.Getwd()
	if len(os.Getenv("MOVIE_SOURCE")) > 0 {
		Source = os.Getenv("MOVIE_SOURCE")
	}

	list := libs.Movie_list(Source)

	for _, path := range list {
		dirname := filepath.Dir(path)

		if len(dirname) == len(Source) {
			continue
		}

		dirname = strings.TrimPrefix(dirname, filepath.Dir(dirname))
		dirname = strings.TrimPrefix(dirname, "/")

		if libs.IsNumber(dirname) {
			newpath := fmt.Sprintf("%s/%s%s", Source, strings.ToUpper(dirname), filepath.Ext(path))
			os.Rename(path, newpath)
		}
	}
}

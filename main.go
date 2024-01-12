package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
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

	for _, sPath := range list {
		dirname := filepath.Dir(sPath)
		sFullFileName := filepath.Base(sPath)
		sFileExt := filepath.Ext(sPath)
		sFileName := strings.TrimSuffix(sFullFileName, sFileExt)
		if len(dirname) == len(Source) {

			// SNIS-009_000^WM_4K.mp4
			re := regexp.MustCompile(`^([A-Za-z]{2,5})-([0-9]{3})_000\SWM_4K`)
			parts := re.FindStringSubmatch(sFileName)
			if len(parts) > 2 {
				sNewFullFileName := fmt.Sprintf("%s-%s-4K-U%s", strings.ToUpper(parts[1]), parts[2], sFileExt)
				sNewPath := filepath.Join(dirname, sNewFullFileName)
				os.Rename(sPath, sNewPath)
				continue
			}

			// ssis-741_000 SSIS-933_000
			re = regexp.MustCompile(`^([A-Za-z]{2,5})-([0-9]{3})_000.*`)
			parts = re.FindStringSubmatch(sFileName)
			if len(parts) > 2 {
				sNewFullFileName := fmt.Sprintf("%s-%s-U%s", strings.ToUpper(parts[1]), parts[2], sFileExt)
				sNewPath := filepath.Join(dirname, sNewFullFileName)
				os.Rename(sPath, sNewPath)
				continue
			}

			// ppsd00046hhb_000^WM mide00612_000^WM
			re = regexp.MustCompile(`^([A-Za-z]{2,5})00([0-9]{3})(hhb)?_000.*`)
			parts = re.FindStringSubmatch(sFileName)
			if len(parts) > 2 {
				sNewFullFileName := fmt.Sprintf("%s-%s-U%s", strings.ToUpper(parts[1]), parts[2], sFileExt)
				sNewPath := filepath.Join(dirname, sNewFullFileName)
				os.Rename(sPath, sNewPath)
				continue
			}

			// PXH-016-4k_000^WM
			re = regexp.MustCompile(`^([A-Za-z]{2,5})-([0-9]{3})-4k_000.*`)
			parts = re.FindStringSubmatch(sFileName)
			if len(parts) > 2 {
				sNewFullFileName := fmt.Sprintf("%s-%s-4K-U%s", strings.ToUpper(parts[1]), parts[2], sFileExt)
				sNewPath := filepath.Join(dirname, sNewFullFileName)
				os.Rename(sPath, sNewPath)
				continue
			}

			// ONED-820_1_000^WM ONED-820_2_000^WM
			re = regexp.MustCompile(`^([A-Za-z]{2,5})-([0-9]{3})_([1-6]{1})_000.*`)
			parts = re.FindStringSubmatch(sFileName)
			if len(parts) > 3 {
				sNewFullFileName := fmt.Sprintf("%s-%s_%s-U%s", strings.ToUpper(parts[1]), parts[2], parts[3], sFileExt)
				sNewPath := filepath.Join(dirname, sNewFullFileName)
				os.Rename(sPath, sNewPath)
				continue
			}
			continue
		}

		dirname = strings.TrimPrefix(dirname, filepath.Dir(dirname))
		dirname = strings.TrimPrefix(dirname, "/")

		if libs.IsNumber(dirname) {
			newpath := fmt.Sprintf("%s/%s%s", Source, strings.ToUpper(dirname), filepath.Ext(sPath))
			os.Rename(sPath, newpath)
		}
	}
}

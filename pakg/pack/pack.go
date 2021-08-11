package pack

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var excludeDir []string

func Zip(src string, dst string, exclude string) error {
	destinationFile, err := os.Create(dst)
	defer destinationFile.Close()

	if err != nil {
		return err
	}
	myZip := zip.NewWriter(destinationFile)
	defer myZip.Close()
	if exclude != "" {
		excludeDir = strings.Split(exclude, ",")
	}

	err = filepath.Walk(src, func(filePath string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if err != nil {
			return err
		}
		if excludeDir != nil {
			for _, dir := range excludeDir {
				if strings.Contains(filePath, dir) {
					return nil
				}
			}
		}
		relPath := strings.TrimPrefix(filePath, filepath.Dir(src))
		zipFile, err := myZip.Create(relPath)
		if err != nil {
			return err
		}
		fsFile, err := os.Open(filePath)
		defer fsFile.Close()
		if err != nil {
			return err
		}
		_, err = io.Copy(zipFile, fsFile)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

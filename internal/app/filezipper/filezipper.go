package filezipper

import (
	"archive/zip"
	"fmt"
	"github.com/inferal73/filezipper/internal/app/logger"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var l = logger.GetLogger()

func Zip(config *Config) error {
	err := l.Log("Start processing\n")
	if err != nil {
		return err
	}

	defer func() {
		err = l.Log("Finished")
	}()

	info, err := os.Stat(config.Entry)
	if err != nil {
		return err
	}

	if info.IsDir() {
		err = zipDir(config.Entry, config.Out)
		if err != nil {
			return err
		}
		return nil
	}
	err = zipFile(config.Entry, config.Out)
	if err != nil {
		return err
	}
	return nil
}

func zipDir(entry, out string) error {
	err := filepath.Walk(entry, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		err = zipFile(path, out)
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

func zipFile(entry, out string) error {
	// get entry file info
	entryFile, err := os.Stat(entry)
	if err != nil {
		return err
	}

	zipFile, err := generateFile(out, entryFile.Name())
	if err != nil {
		return err
	}
	defer func() {
		err = zipFile.Close()
	}()

	// init zip file writer
	archive := zip.NewWriter(zipFile)
	if err != nil {
		return err
	}
	defer func() {
		err = archive.Close()
	}()

	// create zip file headers
	header, err := zip.FileInfoHeader(entryFile)
	if err != nil {
		return err
	}
	header.Method = zip.Deflate

	// open entry file
	zFile, err := os.Open(entry)
	defer func() {
		err = zFile.Close()
	}()
	if err != nil {
		return err
	}
	writer, err := archive.CreateHeader(header)
	if err != nil {
		return err
	}

	// compress file
	_, err = io.Copy(writer, zFile)
	if err != nil {
		return err
	}
	return err
}

func generateFile(out, fileName string) (*os.File, error) {
	zipFileName := changeFileExt(fileName, "zip")
	err := l.Log("Archiving %s ...\n", zipFileName)
	if err != nil {
		return nil, err
	}
	zipFile, _, err := createFile(out, zipFileName)
	if err != nil {
		return nil, err
	}
	return zipFile, nil
}

func changeFileExt(fileName string, ext string) string {
	fileExt := filepath.Ext(fileName)
	withOutExt := strings.TrimSuffix(fileName, fileExt)
	return fmt.Sprintf("%s.%s", withOutExt, ext)
}

func createFile(path string, fileName string) (*os.File, string, error) {
	_, err := os.Stat(path)
	if err != nil {
		if !os.IsNotExist(err) {
			return nil, "", err
		}
		err = os.MkdirAll(path, 0750)
		if err != nil {
			return nil, "", err
		}
	}
	pathWithFile := filepath.Join(path, fileName)
	file, err := os.Create(pathWithFile)
	if err != nil {
		return nil, "", err
	}
	return file, pathWithFile, nil
}
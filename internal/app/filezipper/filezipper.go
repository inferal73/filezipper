package filezipper

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func changeFileExt(fileName string, ext string) string {
	fileExt := filepath.Ext(fileName)
	withOutExt := strings.TrimSuffix(fileName, fileExt)
	return fmt.Sprintf("%s.%s", withOutExt, ext)
}

func createZipFile(path string, fileName string) (*os.File, string, error) {
	if _, err := os.Stat(path);
		os.IsNotExist(err) {
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

func ZipFiles(config *Config, w io.Writer) error {
	_, err := fmt.Fprintf(w, "Start processing\n")
	if err != nil {
		return err
	}

	defer func() {
		_, err = fmt.Fprintf(w,"Finished")
	}()

	info, err := os.Stat(config.Entry)
	if err != nil {
		return err
	}

	if !info.IsDir() {
		_, err = ZipFile(config.Entry, config.Out, w)
		if err != nil {
			return err
		}
		return nil
	}

	err = filepath.Walk(config.Entry, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		_, err = ZipFile(path, config.Out, w)
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

func ZipFile(entry, out string, w io.Writer) (path string, err error) {
	// get entry file info
	entryFile, err := os.Stat(entry)
	if err != nil {
		return
	}

	// generate zip file name
	zipFileName := changeFileExt(entryFile.Name(), "zip")
	_, err = fmt.Fprintf(w, "Archiving %s ...\n", zipFileName)
	if err != nil {
		return
	}

	// create zip file
	zipFile, path, err := createZipFile(out, zipFileName)
	if err != nil {
		return
	}
	defer func() {
		err = zipFile.Close()
	}()

	// init zip file writer
	archive := zip.NewWriter(zipFile)
	defer archive.Close()

	// create zip file headers
	header, err := zip.FileInfoHeader(entryFile)
	if err != nil {
		return
	}
	header.Method = zip.Deflate

	// open entry file
	zFile, err := os.Open(entry)
	defer func() {
		err = zFile.Close()
	}()
	if err != nil {
		return
	}
	writer, err := archive.CreateHeader(header)
	if err != nil {
		return
	}

	// compress file
	_, err = io.Copy(writer, zFile)
	if err != nil {
		return
	}
	return
}
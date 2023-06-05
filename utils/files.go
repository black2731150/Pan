package utils

import (
	"io/fs"
	"os"
	"pan/models"
	"path/filepath"
	"strings"
)

func MakeDir(path string) error {
	err := os.Mkdir(path, 0755)
	if err != nil {
		return err
	}
	return nil
}

func MakeFile(path string, filename string) error {
	file, err := os.Create(path + "/" + filename)
	if err != nil {
		return err
	}
	defer file.Close()
	return err
}

func OpenFile(filepath string) (*os.File, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func GetFileStat(filepath string) (os.FileInfo, error) {
	file, err := OpenFile(filepath)
	if err != nil {
		return nil, err
	}

	filestat, err := file.Stat()
	if err != nil {
		return nil, err
	}

	return filestat, nil
}

func RenameFile(filepath string, newname string) error {
	lastindex := strings.LastIndex(filepath, "/")
	newpath := filepath[:lastindex] + "/" + newname
	err := os.Rename(filepath, newpath)
	if err != nil {
		return err
	}
	return nil
}

func RemoveFile(filepath string) error {
	err := os.Remove(filepath)
	if err != nil {
		return err
	}
	return nil
}

func GetFilesInfoFromFolder(folderpath string) ([]models.Fileinfo, error) {
	fileinfo := []models.Fileinfo{}
	err := filepath.Walk(folderpath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if path == folderpath {
			return nil
		}

		if info.IsDir() {
			var onefileinfo models.Fileinfo
			onefileinfo.Name = info.Name()
			onefileinfo.Size = info.Size()
			onefileinfo.IsDir = info.IsDir()
			onefileinfo.Perm = info.Mode().Perm()
			onefileinfo.ModTime = info.ModTime()
			fileinfo = append(fileinfo, onefileinfo)
			return filepath.SkipDir
		} else {
			var onefileinfo models.Fileinfo
			onefileinfo.Name = info.Name()
			onefileinfo.Size = info.Size()
			onefileinfo.IsDir = info.IsDir()
			onefileinfo.Perm = info.Mode().Perm()
			onefileinfo.ModTime = info.ModTime()
			fileinfo = append(fileinfo, onefileinfo)
			return nil
		}
	})

	if err != nil {
		return nil, err
	}

	return fileinfo, nil
}

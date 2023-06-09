package utils

import (
	"fmt"
	"io/fs"
	"os"
	"pan/models"
	"path/filepath"
)

func MakeDir(basePath string) error {
	directory := basePath
	counter := 1

	for {
		_, err := os.Stat(directory)
		if os.IsNotExist(err) {
			// 目录不存在，创建新目录
			err := os.Mkdir(directory, 0755)
			if err != nil {
				return err
			}
			return nil
		} else if err != nil {
			// 发生其他错误
			return err
		}

		// 目录已存在，尝试创建下一个唯一目录
		directory = fmt.Sprintf("%s(%d)", basePath, counter)
		counter++
	}
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

func RenameFile(filepath string, newpath string) error {
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

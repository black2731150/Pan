package models

import (
	"io/fs"
	"time"
)

type Fileinfo struct {
	Name    string      //文件名称
	Size    int64       //文件大小
	IsDir   bool        //是否是目录
	ModTime time.Time   //最后修改时间
	Perm    fs.FileMode //文件权限
}

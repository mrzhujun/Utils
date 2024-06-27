package file

import (
	"path/filepath"
	"runtime"
)

func GetMyDir() string {
	_, filename, _, _ := runtime.Caller(0)
	// 获取当前文件的绝对路径
	absPath, _ := filepath.Abs(filename)
	// 获取当前文件所在的目录
	return filepath.Dir(absPath)

}

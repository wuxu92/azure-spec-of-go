package utils

import (
	"path"
	"runtime"
)

var pwd string

func init() {
	_, file, _, _ := runtime.Caller(0)
	pwd = path.Dir(path.Dir(file))
}

func ProjectDir() string {
	return pwd
}

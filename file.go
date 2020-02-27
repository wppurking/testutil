package testutil

import (
	"io/ioutil"
	"path/filepath"
	"runtime"
)

var (
	basepath = CallerPath()
)

// CallerPath 返回命令行中第一个命令的调用文件路径
func CallerPath() string {
	_, b, _, _ := runtime.Caller(0)
	return filepath.Dir(b)
}

// ReadTestFile 读取 testFile 到字符串. 读取 testdata 与 test/testdata 目录下的数据
func ReadTestFile(name string) string {
	// TODO 需要测试此方法
	bs, err := ioutil.ReadFile(filepath.Join(basepath, "testdata", name))
	if err != nil {
		panic(err)
	}
	return string(bs)
}

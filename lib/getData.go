package lib

import (
	"io/ioutil"
	. "k8s_yaml/app"
	"os"
)

func GetFile(pwd string) (rootRes, fileRes []os.FileInfo) {
	var (
		FileInfo []os.FileInfo
		FileInfo2 []os.FileInfo
		err error
	)
	// 获取目录
	if FileInfo,err = ioutil.ReadDir("./data"); err != nil {
		Log.Warn(err.Error())
	}

	// 添加根目录列表
	for _,fileInfo := range FileInfo {
		if fileInfo.IsDir() {
			rootRes = append(rootRes, fileInfo)
		}
	}

	// 获取二级目录结构
	if pwd != "" {
		if FileInfo2,err = ioutil.ReadDir("./data/"+pwd); err != nil {
			Log.Warn(err.Error())
		}
		for _,fileInfo := range FileInfo2 {
			if !fileInfo.IsDir() {
				fileRes = append(fileRes, fileInfo)
			}
		}
	}
	return
}
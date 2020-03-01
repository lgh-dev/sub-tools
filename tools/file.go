package tools

import (
	"fmt"
	"os"
)

// 递归创建不存在的目录。
func MkdirIfNotExists(path string) {
	exist, err := PathExists(path)
	if err != nil {
		fmt.Printf("get dir error![%v]\n", err)
		return
	}
	if !exist {
		fmt.Printf("no dir![%v]\n", path)
		// 创建文件夹
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			fmt.Printf("mkdir failed![%v]\n", err)
		} else {
			fmt.Printf("mkdir success!\n")
		}
	}
}

// 判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

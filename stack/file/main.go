package main

import (
	"errors"
	"fmt"
	"io/ioutil"
)

func main() {
	path := "/Users/shuwen/go/src/shuwen-algorithm"
	files := []string{}
	files, _ = GetAll(path, files)
	for i := 0; i < len(files); i++ {
		fmt.Println("file = ", files[i])
	}
}

// GetAll Get all file
func GetAll(path string, files []string) ([]string, error) {
	read, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, errors.New("File dir don't read!")
	}
	for _, fi := range read {
		if fi.IsDir() { // 判断是否是文件夹
			fullDir := path + "/" + fi.Name()   // 构造新的路径
			files := append(files, fullDir)     // 追加路径
			files, err = GetAll(fullDir, files) // 文件夹递归调用
			if err != nil {
				return nil, err
			} else {
				fullDir = path + "/" + fi.Name() // 构造新的路径
				files = append(files, fullDir)
			}
			fmt.Println("files = ", files)
		}
	}
	return files, nil
}

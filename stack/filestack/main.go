package main

import (
	"algorithm/stack/stackarray"
	"errors"
	"fmt"
	"io/ioutil"
)

func main() {
	path := "/Users/shuwen/go/src/shuwen-algorithm"
	files := []string{}
	myStack := stackarray.NewStack()
	myStack.Push(path)
	for !myStack.IsEmpty() {
		// path := myStack.Pop().(string)
		path := myStack.Pop()
		files := append(files, path)
		fileInfoes, _ := ioutil.ReadDir(path) // Read all dirtory in the files
		for _, file := range fileInfoes {
			if file.IsDir() {
				fullDir := path + "/" + file.Name() // 构造新的路径
				myStack.Push(fullDir)
			} else {
				fullDir := path + "/" + file.Name()
				files = append(files, fullDir)
			}
		}
	}
	for i := 0; i < len(files); i++ {
		fmt.Println(files[i])
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

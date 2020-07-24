package main

import (
	"errors"
	"fmt"
	"io/ioutil"
)

func GetAll(path string, files []string) ([]string, error) {
	read, err := ioutil.ReadDir(path)
	if err != nil {
		return files, errors.New("文件不可读")
	}
	for _, fi := range read { // 循环每个文件或者文件夹
		if fi.IsDir() { // 判断是否文件夹
			fullDir := path + "/" + fi.Name() // 构造新的路径
			files = append(files, fullDir)    // 追加路径
			files, _ = GetAll(fullDir, files)
		} else {
			fullDir := path + "/" + fi.Name()
			files = append(files, fullDir)
		}
	}
	return files, nil
}

func main() {
	path := "/home/teng/Documents/git/gomachine/"
	files := make([]string, 0)
	files, _ = GetAll(path, files)
	for i := 0; i < len(files); i++ {
		fmt.Println(files[i])
	}

}

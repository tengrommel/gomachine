package main

import (
	"fmt"
	"gomachine/code/ch1/StackArray"
	"io/ioutil"
)

func main() {
	path := "/home/teng/Documents/git/gomachine"
	files := make([]string, 0)
	myStack := StackArray.NewStack()
	myStack.Push(path)
	for !myStack.IsEmpty() {
		path := myStack.Pop().(string)
		files = append(files, path)     // 加入列表
		read, _ := ioutil.ReadDir(path) // 读取文件夹下面所有的路径
		for _, fi := range read {
			if fi.IsDir() {
				fullDir := path + "/" + fi.Name()
				//files = append(files, fullDir)
				myStack.Push(fullDir)
			} else {
				fullDir := path + "/" + fi.Name()
				files = append(files, fullDir)
			}
		}
	}
	for i := 0; i < len(files); i++ {
		fmt.Println(files[i])
	}
}

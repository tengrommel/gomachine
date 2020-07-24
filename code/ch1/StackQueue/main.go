package main

import (
	"errors"
	"fmt"
	"gomachine/code/ch1/Queue"
	"io/ioutil"
)

//
// --
// ----
func GetAll(path string, files []string, level int) ([]string, error) {
	levelStr := ""
	if level == 1 {
		levelStr = "+"
	} else {
		for ; level > 1; level-- {
			levelStr += "|--"
		}
		levelStr += "+"
	}
	read, err := ioutil.ReadDir(path)
	if err != nil {
		return files, errors.New("文件不可读")
	}
	for _, fi := range read { // 循环每个文件或者文件夹
		if fi.IsDir() { // 判断是否文件夹
			fullDir := path + "/" + fi.Name() // 构造新的路径
			files = append(files, fullDir)    // 追加路径
			newLevel := level + 1
			files, _ = GetAll(fullDir, files, newLevel)
		} else {
			fullDir := path + "/" + fi.Name()
			files = append(files, levelStr+fullDir)
		}
	}
	return files, nil
}

func main() {
	//path := "/home/teng/Documents/git/gomachine"
	//files := make([]string, 0)
	//files, _ = GetAll(path, files, 1)
	//for i := 0; i < len(files); i++ {
	//	fmt.Println(files[i])
	//}
	//
	//// 深度
	//myQueue := Queue.NewQueue()
	//myQueue.EnQueue(1)
	//myQueue.EnQueue(2)
	//myQueue.EnQueue(3)
	//myQueue.EnQueue(4)
	//fmt.Println(myQueue.DeQueue())
	//fmt.Println(myQueue.DeQueue())
	//fmt.Println(myQueue.DeQueue())
	//fmt.Println(myQueue.DeQueue())
	//myQueue.EnQueue(14)
	//fmt.Println(myQueue.DeQueue())
	path := "/home/teng/Documents/git/gomachine"
	files := make([]string, 0)
	myStack := Queue.NewQueue()
	myStack.EnQueue(path)
	for {
		path := myStack.DeQueue() // 不断从队列中取出数据
		if path == nil {
			break
		}
		fmt.Println("get", path)
		read, _ := ioutil.ReadDir(path.(string))
		for _, fi := range read {
			if fi.IsDir() {
				fullDir := path.(string) + "/" + fi.Name()
				fmt.Println("Dir", fullDir)
				myStack.EnQueue(fullDir)
			} else {
				fullDir := path.(string) + "/" + fi.Name()
				files = append(files, fullDir)
				fmt.Println("file", fullDir)
			}
		}
	}
	for i := 0; i < len(files); i++ {
		fmt.Println(files[i])
	}
}

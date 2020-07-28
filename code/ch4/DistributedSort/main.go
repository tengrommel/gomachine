package main

import (
	"bufio"
	"fmt"
	"gomachine/code/ch4/DistributedSort/pipeLineMiddleWare"
	"os"
)

// 本地 多线程 分布式
// 生成随机数组
func main() {
	var fileName = "data.in" // 文件写入
	var count = 100000
	file, _ := os.Create(fileName)
	defer file.Close()                               // 延迟关闭文件
	myPipe := pipeLineMiddleWare.RandomSource(count) // 管道装载随机数
	writer := bufio.NewWriter(file)                  // 写入
	pipeLineMiddleWare.WriteSlink(writer, myPipe)    // 写入
	writer.Flush()                                   // 刷新
	file, _ = os.Open(fileName)
	defer file.Close()
	myPipeRead := pipeLineMiddleWare.ReaderSource(bufio.NewReader(file), -1)
	counter := 0
	for v := range myPipeRead {
		fmt.Println(v)
		counter++
		if counter > 1000 {
			break
		}
	}

}

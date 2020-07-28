package main

import (
	"bufio"
	"fmt"
	"gomachine/code/ch4/DistributedSort/single/pipeLineMiddleWare"
	"os"
	"strconv"
	"time"
)

// 本地 多线程 分布式

// 分布式
func createNetworkPipeLine(fileName string, fileSize int, chunkCount int) <-chan int {
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	myPipe := pipeLineMiddleWare.RandomSource(fileSize / 8)
	writer := bufio.NewWriter(file)
	pipeLineMiddleWare.WriteSlink(writer, myPipe)
	writer.Flush()

	chunkSize := fileSize / chunkCount
	sortAddr := make([]string, 0)
	pipeLineMiddleWare.Init()
	file, err = os.Open(fileName)
	if err != nil {
		panic(err)
	}
	for i := 0; i < chunkCount; i++ {
		file.Seek(int64(i*chunkSize), 0) // 移动文件指针位置
		source := pipeLineMiddleWare.ReaderSource(bufio.NewReader(file), chunkSize)
		addr := ":" + strconv.Itoa(7000+i) // 开辟地址
		pipeLineMiddleWare.NetWordWrite(addr, pipeLineMiddleWare.InMemorySort(source))
		sortAddr = append(sortAddr, addr) // 地址复制
	}
	sortResults := make([]<-chan int, 0)
	for _, addr := range sortAddr {
		sortResults = append(sortResults, pipeLineMiddleWare.NetWordRead(addr))
	}
	return pipeLineMiddleWare.MergeN(sortResults...)
}

// 多线程 - 调用中间件完成
func createPipeLine(fileName string, fileSize int, chunkCount int) <-chan int {
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	myPipe := pipeLineMiddleWare.RandomSource(fileSize / 8)
	writer := bufio.NewWriter(file)
	pipeLineMiddleWare.WriteSlink(writer, myPipe)
	writer.Flush()

	chunkSize := fileSize / chunkCount   // 数量
	sortResults := make([]<-chan int, 0) // 排序结果，一个数组，每一个元素是一个管道
	pipeLineMiddleWare.Init()            // 初始化
	for i := 0; i < chunkCount; i++ {
		file, err = os.Open(fileName) // 打开文件
		if err == nil {
			panic(err)
		}
		file.Seek(int64(i*chunkSize), 0)                                            // 跳到文件指针
		source := pipeLineMiddleWare.ReaderSource(bufio.NewReader(file), chunkSize) // 读取
		sortResults = append(sortResults, pipeLineMiddleWare.InMemorySort(source))  // 结果排序
	}
	return pipeLineMiddleWare.MergeN(sortResults...)
}

// 写入文件
func writeToFile(in <-chan int, fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	defer writer.Flush()                      // 刷新
	pipeLineMiddleWare.WriteSlink(writer, in) // 写入数据
}

// 显示文件
func showFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	p := pipeLineMiddleWare.ReaderSource(bufio.NewReader(file), -1)
	counter := 0
	for v := range p {
		fmt.Println(v)
		counter++
		if counter > 1000 {
			break
		}
	}
}

// 生成随机数组
func main() {
	//var fileName = "data.in" // 文件写入
	//var count = 100000
	//file, _ := os.Create(fileName)
	//defer file.Close()                               // 延迟关闭文件
	//myPipe := pipeLineMiddleWare.RandomSource(count) // 管道装载随机数
	//writer := bufio.NewWriter(file)                  // 写入
	//pipeLineMiddleWare.WriteSlink(writer, myPipe)    // 写入
	//writer.Flush()                                   // 刷新
	//file, _ = os.Open(fileName)
	//defer file.Close()
	//myPipeRead := pipeLineMiddleWare.ReaderSource(bufio.NewReader(file), -1)
	//counter := 0
	//for v := range myPipeRead {
	//	fmt.Println(v)
	//	counter++
	//	if counter > 1000 {
	//		break
	//	}
	//}
	go func() {
		p := createPipeLine("big.in", 800000, 4)
		writeToFile(p, "big.out")
		showFile("big.out")
	}()
	time.Sleep(1 * time.Second)
}

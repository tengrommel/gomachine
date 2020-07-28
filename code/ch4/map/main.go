package main

import (
	"fmt"
	"sync"
	"time"
)

// map映射
// map管理上亿的数据，瞬间查找

type SyncMap struct {
	myMap map[string]string
	lock  *sync.RWMutex // 读写锁
}

var sMap SyncMap   // 公有的map
var done chan bool // 通道是否完成

func write1() {
	keys := []string{"1", "2", "3"}
	for _, k := range keys {
		sMap.lock.Lock()
		sMap.myMap[k] = k // 赋值
		sMap.lock.Unlock()
		time.Sleep(1 * time.Second)
	}
	done <- true // 通道写入我们干完了
}

func write2() {
	keys := []string{"a1", "b2", "c3"}
	for _, k := range keys {
		sMap.lock.Lock()
		sMap.myMap[k] = k // 赋值
		sMap.lock.Unlock()
		time.Sleep(1 * time.Second)
	}
	done <- true // 通道写入我们干完了
}

func read() {
	sMap.lock.RLock() // 读的时候加锁
	fmt.Println("readLock")
	for k, v := range sMap.myMap {
		fmt.Println(k, v)
	}
	sMap.lock.RUnlock()
}

func main() {
	sMap = SyncMap{make(map[string]string), new(sync.RWMutex)}
	done = make(chan bool, 2) // 管道处理写入
	go write1()
	go write2()
	for {
		read()
		if len(done) == 2 {
			fmt.Println(sMap.myMap)
			break
		} else {
			time.Sleep(1 * time.Second) // 没有完成继续等待
		}
	}
}

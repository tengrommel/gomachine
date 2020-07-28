package main

import (
	"fmt"
	"sync"
	"time"
)

var money int
var lock *sync.RWMutex = new(sync.RWMutex)

func add(pint *int) {
	lock.Lock()
	for i := 0; i < 100000; i++ {
		*pint++
	}
	lock.Unlock()
}

func main() {
	for i := 0; i < 1000; i++ {
		go add(&money)
	}
	time.Sleep(time.Second * 5)
	fmt.Println(money)
}

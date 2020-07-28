package main

import "time"

var money int

func add(pint *int) {
	for i := 0; i < 100000; i++ {
		*pint++
	}
}

func main() {
	for i := 0; i < 1000; i++ {
		go add(&money)
	}
	time.Sleep(time.Second * 20)
}

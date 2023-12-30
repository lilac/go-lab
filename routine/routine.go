package main

import (
	"fmt"
	"log"
	"os"
	"runtime/trace"
	"sync"
	"time"
)

var out = make(chan int)
var wg sync.WaitGroup

func worker(num int) {
	defer wg.Done()
	time.Sleep(time.Second)
	out <- num * 2
}

func main() {
	// 创建trace文件
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// 启动trace goroutine
	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	routines := 10000
	defer duration(track("Real time"))
	for i := 0; i < routines; i++ {
		wg.Add(1)
		go worker(i)
	}
	//wg.Wait()
	count := 0
	for i := 0; i < routines; i++ {
		<-out
		count++
	}
	fmt.Printf("Result count: %v\n", count)
}

func track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func duration(msg string, start time.Time) {
	log.Printf("%v: %v\n", msg, time.Since(start))
}

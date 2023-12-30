package main

import (
	"fmt"
	"github.com/arl/statsviz"
	"log"
	"net/http"
	"syscall"
	"time"
)

const pages = 1000

func main() {
	fmt.Printf("Page size: %v\n", syscall.Getpagesize())
	fmt.Printf("PID: %v\n", syscall.Getpid())
	mux := http.NewServeMux()
	statsviz.Register(mux)
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", mux))
	}()
	log.Println("Start stressing.")
	streeCPU(2)
	stressMemory(pages, time.Minute*5)
}

func handler(w http.ResponseWriter, r *http.Request) {
}

func streeCPU(cores int) {
	for i := 0; i < cores; i++ {
		go func() {
			for {
			}
		}()
	}
}
func stressMemory(pages int, duration time.Duration) {
	b, err := syscall.Mmap(-1, 0, pages*syscall.Getpagesize(), syscall.PROT_WRITE, syscall.MAP_ANON|syscall.MAP_PRIVATE)
	if err != nil {
		fmt.Printf("Mmap: %v", err)
	} else {
		defer func() {
			if err := syscall.Munmap(b); err != nil {
				fmt.Printf("Munmap: %v", err)
			}
		}()
		for i := 0; i < pages; i++ {
			index := i * syscall.Getpagesize()
			b[index] = '0'
		}
		time.Sleep(duration)
	}
}

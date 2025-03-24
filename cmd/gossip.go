package main

import (
	"fmt"
	"sync"

	webrtc "github.com/14253647/gossip/pkg/webrtc"
)

func worker(id int, task func(), wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("worker %d starting\n", id)
	task()
	fmt.Printf("worker %d done\n", id)
}

func main() {
	fmt.Println("Starting cmd/gossip")
	defer fmt.Println("Stopped cmd/gossip")

	var wg sync.WaitGroup
	worker(0, func() { webrtc.RunServer("stun:stun.l.google.com:19302") }, &wg)
	wg.Wait()
}

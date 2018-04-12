package main

import (
	"fmt"
	"math/rand"
	"time"
)

func SendRandomInt(sendCh chan<- int) {

	randInt := rand.Int63n(60)
	time.Sleep(time.Second * time.Duration(randInt))

	sendCh <- rand.Int()
}

func main() {

	// not buffered so sending to channel requires go routines or deadlock would happen
	intCh := make(chan int)

	for i := 0; i < 5; i++ {
		go SendRandomInt(intCh)
	}

	for i := 0; i < 5; i++ {
		go fmt.Println(<-intCh)
	}

	// buffered

	intChBuffered := make(chan int, 5)

	for i := 0; i < 5; i++ {
		SendRandomInt(intChBuffered)
	}

	for i := 0; i < 5; i++ {
		fmt.Println(<-intChBuffered)
	}

}

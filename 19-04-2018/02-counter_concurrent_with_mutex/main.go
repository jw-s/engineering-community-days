package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func AddToCounter(counterCh chan uint64, index int) {
	counter := <-counterCh

	fmt.Printf("I am worker %v and the counter is currently at: %v\n", index, counter)

	randAmountToAdd := rand.Uint64()
	fmt.Printf("I am worker %v and will increment by %v\n", index, randAmountToAdd)

	counter += randAmountToAdd

	counterCh <- counter
}

func main() {

	var wg sync.WaitGroup

	counterCh := make(chan uint64, 1)

	fmt.Println("Sending seed value")

	counterCh <- 0

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			AddToCounter(counterCh, index)

		}(i)
	}

	wg.Wait()

	fmt.Println(<-counterCh)

}

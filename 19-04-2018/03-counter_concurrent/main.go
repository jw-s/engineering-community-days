package main

import (
	"fmt"
	"math/rand"
	"time"
)

func SendToBeAddedToCounter(counterCh chan<- int32, index int) {

	randAmountToAdd := rand.Int31n(50)
	fmt.Printf("I am worker %v and will add %v to the counter \n", index, randAmountToAdd)

	counterCh <- randAmountToAdd
}

func accumulator(counter <-chan int32, done <-chan struct{}) int32 {

	var cnt int32

	for {
		select {
		case i := <-counter:

			cnt += i

		case <-done:

			return cnt
		}
	}
}

func main() {

	counterCh := make(chan int32) //channel which will read from senders and

	doneCh := make(chan struct{}) //channel to signal accumulator to return the final count

	for i := 0; i < rand.Intn(1000000); i++ { //spawn random amount of "workers" to send ints to be added to the counter
		go func(index int) {
			SendToBeAddedToCounter(counterCh, index)
		}(i)
	}

	go func(done chan<- struct{}) { // spawn cancel function which sends a signal to return final counter value
		randTime := time.Second * time.Duration(rand.Int63n(60))

		fmt.Printf("Stopping accumulation in %v seconds\n", randTime)

		time.Sleep(randTime)

		close(doneCh) //empty struct is zero memory allocation and idiomatic in go
	}(doneCh)

	fmt.Println(accumulator(counterCh, doneCh)) //consumes incomming counter additons send from workers and adds to existing counter

}

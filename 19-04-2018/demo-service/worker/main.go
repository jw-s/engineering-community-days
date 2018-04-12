package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

func DispatchWork(workers chan chan struct{}, doneCh chan struct{}) {
	for {
		select {
		case worker := <-workers:
			worker <- struct{}{}

		case <-doneCh:

			return

		}
	}
}

func GetRandomFact() (string, error) {
	randomTimeToSleep := rand.Intn(60)

	res, err := http.Get("https://random-quote-generator.herokuapp.com/api/quotes/random")

	if err != nil {
		return "", nil
	}

	b, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return "", nil
	}

	time.Sleep(time.Second * time.Duration(randomTimeToSleep))

	return string(b), nil
}

func main() {

	workers := make(chan chan struct{}, 5)

	doneCh := make(chan struct{})

	go DispatchWork(workers, doneCh)

	for i := 0; i < 5; i++ {
		go func(askForWork chan<- chan struct{}, done <-chan struct{}, index int) {
			myWorkCh := make(chan struct{})
			askForWork <- myWorkCh
			for {
				select {
				case <-myWorkCh:
					fmt.Printf("I've been asked to work, I am worker %v\n", index)

					fact, err := GetRandomFact()

					if err != nil {
						fmt.Println("exiting due to error")
						return
					}
					fmt.Printf("Worker %v presents you with the fact: %s\n", index, fact)

					fmt.Printf("finished work: worker %v, ready to pick up more work\n", index)
					askForWork <- myWorkCh
				case <-done:
					fmt.Println("exiting")
					return
				}
			}
		}(workers, doneCh, i)
	}

	time.Sleep(time.Minute * 5)

	close(doneCh)

}

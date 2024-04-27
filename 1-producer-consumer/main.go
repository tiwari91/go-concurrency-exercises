package main

import (
	"fmt"
	"sync"
	"time"
)

func producer(stream Stream, tweetsCh chan<- *Tweet, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(tweetsCh)
	for {
		tweet, err := stream.Next()
		if err == ErrEOF {
			return
		}
		tweetsCh <- tweet
	}
}

func consumer(tweetsCh <-chan *Tweet, wg *sync.WaitGroup) {
	defer wg.Done()
	for t := range tweetsCh {
		if t.IsTalkingAboutGo() {
			fmt.Println(t.Username, "\ttweets about golang")
		} else {
			fmt.Println(t.Username, "\tdoes not tweet about golang")
		}
	}
}

func main() {
	start := time.Now()
	stream := GetMockStream()

	var wg sync.WaitGroup

	tweetsCh := make(chan *Tweet)
	wg.Add(2)

	go producer(stream, tweetsCh, &wg)

	go consumer(tweetsCh, &wg)

	wg.Wait()

	fmt.Printf("Process took %s\n", time.Since(start))
}

//////////////////////////////////////////////////////////////////////
//
// Given is a producer-consumer scenario, where a producer reads in
// tweets from a mockstream and a consumer is processing the
// data. Your task is to change the code so that the producer as well
// as the consumer can run concurrently
//

package main

import (
	"fmt"
	"time"
)

func producer(stream Stream, channel chan *Tweet) (tweets []*Tweet) {
	for {
		tweet, err := stream.Next()
		if err == ErrEOF {
			close(channel)
			return tweets
		}

		channel <- tweet
	}
}

func consumer(tweets []*Tweet, channel chan *Tweet) {
	for t := range channel {
		go func(t *Tweet) {
			if t.IsTalkingAboutGo() {
				fmt.Println(t.Username, "\ttweets about golang")
			} else {
				fmt.Println(t.Username, "\tdoes not tweet about golang")
			}
		}(t)
	}
}

func main() {
	start := time.Now()
	stream := GetMockStream()
	channel := make(chan *Tweet, 100)

	// Producer
	tweets := producer(stream, channel)

	// Consumer
	consumer(tweets, channel)

	fmt.Printf("Process took %s\n", time.Since(start))
}

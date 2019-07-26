package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"sync"
)

func main() {
	testNSQ()
}

type NSQHandler struct {
}

func (this *NSQHandler) HandleMessage(msg *nsq.Message) error {
	fmt.Println("receive", msg.NSQDAddress, "message:", string(msg.Body))
	return nil
}

func testNSQ() {

	url := "127.0.0.1:4150"
	wait := sync.WaitGroup{}
	wait.Add(1)

	go func() {
		defer wait.Done()
		config := nsq.NewConfig()
		config.MaxInFlight = 9

		for i := 0; i < 10; i++ {
			consumer, err := nsq.NewConsumer("test", "one", config)
			if nil != err {
				fmt.Println("err", err)
				return
			}

			consumer.AddHandler(&NSQHandler{})
			err = consumer.ConnectToNSQD(url)
			if nil != err {
				fmt.Println("err", err)
				return
			}
		}
		select {}
	}()

	wait.Wait()
}

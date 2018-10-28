package main

import (
	"fmt"
	"github.com/segmentio/kafka-go"
	"context"
	"sync"
	"github.com/kataras/go-events"
)
var wg sync.WaitGroup


func (k *kel) Sub()  {
	wg.Add(len(k.Channels))
	//multi thread with gorutes
	for i := 0; i < len(k.Channels); i++ {
		go  k.subEventTopic(k.Channels[i])
	}
	wg.Wait()
	fmt.Println("Finished for loop")
}


func (k *kel) subEventTopic(topic string){

	fmt.Println("Channel :: ", topic)
	//Init config for kafka lib
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{k.Adress},
		GroupID:   k.EName,
		Topic:     topic,
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
	})

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}
		//Fire event (github.com/kataras/go-events)
		events.Emit(events.EventName(topic), m.Topic, m.Offset, m.Partition, string(m.Key), string(m.Value))
	}
	r.Close()
	defer wg.Done()

}
package main

import (
	"fmt"
	"encoding/json"
	"context"
	"time"
	"github.com/segmentio/kafka-go"
)


type Event struct {
	//Event :: the basic event, needed for the encoder.
	U string `json:"u"` //UUID from the user Acc (#NO)
	E string `json:"e"` // Event name (channel or topic)
	D interface{} `json:"d"` // The data object.
}


func (k *kel)SendEventToChannel(cchannel string, event interface{})  {
	allowed := false;
	for i := 0; i < len(k.Channels); i++ {
		if (k.Channels[i] == cchannel){
			allowed = true;


			eventS := Event{D:event, E:cchannel, U: "NO",}

			b, err := json.Marshal(eventS)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("-> " + string(b))

			conn, _ := kafka.DialLeader(context.Background(), "tcp", k.Adress, cchannel, 0)

			conn.SetWriteDeadline(time.Now().Add(10*time.Second))
			conn.WriteMessages(
				kafka.Message{Value: []byte(b)},
			)

			conn.Close()

		}

	}
	if (!allowed){
		fmt.Println("[ERROR] You try to send a event to a not allowed channel! Permission denied!")
	}

}
package main

import (
	"kafkaEventLib/ct"
	"github.com/kataras/go-events"
	"fmt"
)

func main() {



	rf := kel{
		Channels: []string{ct.UserBlocked, ct.UserLogedin},
	}
	go rf.Init("34.247.57.14:9092", "m-endpoint-login")


	events.On(ct.UserBlocked, func(msg ...interface{}) {
		/*		Kafka Message
			   [0] = Topic
			   [1] = Offset
			   [2] = Partition (get only the sub Partition)
			   [3] = Key
			   [4] = Value (EVENT DATA)
		*/

		fmt.Println("TRIGGERED!!!!!!!!! :: ", msg[4])
		type Login struct {
			//Event :: the basic event, needed for the encoder.
			Username string `json:"username"` //UUID from the user Acc (#NO)
			Email string `json:"email"` // Event name (channel or topic)

		}

		l:= Login{
			Email: "test@test.de",
			Username: "VerHext",
		}

		rf.SendEventToChannel(ct.UserLogedin, l  )

	})

	type Login struct {
		//Event :: the basic event, needed for the encoder.
		Username string `json:"username"` //UUID from the user Acc (#NO)
		Email string `json:"email"` // Event name (channel or topic)

	}

	l:= Login{
		Email: "test@test.de",
		Username: "VerHext",
	}

	rf.SendEventToChannel(ct.UserLogedin, l  )

	rf.Sub();
}




# kafka-event-lib

This is a lib designed by the magiCCoder organisation. That help use to developed a high scaled ans secure API based on the kafka protocol. 
We implement:
  - event handler
  - permission system
  - event sender
  
# Wy use we a lib?
Because, we want to developed lots of small "micro API", and we don't want to implement in echa API the same code. So we use one lib. That help us in maintaince...

#Example
The lib is simple to use, there is a exampe:

```go
package main

import (
	"github.com/kafkaEventLib/ct"
	"github.com/kataras/go-events"
	"fmt"
)

func main() {

	rf := kel{
		Channels: []string{ct.UserBlocked, ct.UserLogedin},
	}
                                //Thats the consumer name
	go rf.Init("127.0.0.1:9092", "m-endpoint-login")


	events.On(ct.UserBlocked, func(msg ...interface{}) {
		/*		Kafka Message
			   [0] = Topic
			   [1] = Offset
			   [2] = Partition (get only the sub Partition)
			   [3] = Key
			   [4] = Value (EVENT DATA)
		*/

		fmt.Println("[EVENT] :: ", msg[4])
		type Login struct {
			//Event :: the basic event, needed for the encoder.
			Username string `json:"username"` //UUID from the user Acc (#NO)
			Email string `json:"email"` // Event name (channel or topic)

		}

		l:= Login{ Email: "test@test.de", Username: "VerHext",}
		rf.SendEventToChannel(ct.UserLogedin, l  )

	})
//Start the listener
	rf.Sub();
}

```

# How can I use the lib?
This is not a api, thats are developed for any workflow. 
Please set your own Topics in the "enum" class.  
And if you want to use a other Event design, you should change it.

_______________________
# Developed by

![](https://sounds.support-pp.de/magiccoder-brand.png)

package nats_server

import (
	"encoding/json"
	"fmt"
	"time"
	"user/api/dto"

	"github.com/nats-io/nats.go"
)

func Subscriber(subject string) {
	time.Sleep(time.Second * 5)

	nc, err := nats.Connect("nats")
	if err != nil {
		panic(err)
	}

	received := &dto.SendMessageRequest{}

	go func() {
		not_have := 0
		for {
			if len(received.Message) > 1 {
				fmt.Printf("\n\nUSER:: MESSAGE:\t%s\n\n", received.Message)
				return
			} else {
				fmt.Printf("\n\nUSER:: Not have anything yet #%d\n\n", not_have)
				not_have++
			}
			time.Sleep(time.Second * 5)
		}
	}()

	nc.Subscribe(subject, func(m *nats.Msg) {
		err := json.Unmarshal([]byte(string(m.Data)), &received)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("\n~~~~~~~~~~~~~~~~~~~~~~~\nUSER:: Received a request ...\nChatName:\t%s\nFrom:\t\t%s\nTo:\t\t%s\nMessage:\t%s\n~~~~~~~~~~~~~~~~~~~~~~~\n", received.ChatName, received.From, received.To, received.Message)
		nc.Publish(fmt.Sprintf("%s_response", subject), []byte("I got ya"))
		return
	})

}

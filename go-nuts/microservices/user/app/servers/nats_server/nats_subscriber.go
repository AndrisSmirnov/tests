package nats_server

import (
	"encoding/json"
	"fmt"
	"time"
	"user/api/dto"

	"github.com/nats-io/nats.go"
)

func Subscriber(subject string) {
	nc, err := nats.Connect("nats")
	if err != nil {
		panic(err)
	}
	fmt.Println(nc)

	received := &dto.SendMessageRequest{}

	go func() {
		not_have := 0
		for {
			if len(received.Message) > 1 {
				fmt.Printf("\n\nMESSAGE:\t%s\n\n", received.Message)
				return
			} else {
				fmt.Printf("\n\nNot have anything yet #%d\n\n", not_have)
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

		fmt.Printf("\n~~~~~~~~~~~~~~~~~~~~~~~\nReceived a request ...\nChatName:\t%s\nFrom:\t\t%s\nTo:\t\t%s\nMessage:\t%s\n~~~~~~~~~~~~~~~~~~~~~~~\n", received.ChatName, received.From, received.To, received.Message)
		return
	})

}

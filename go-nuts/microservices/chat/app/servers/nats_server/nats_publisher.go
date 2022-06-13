package nats_server

import (
	dto "chat/api/dto"
	"fmt"
	"time"

	nats "github.com/nats-io/nats.go"
)

func Publisher(subject, From, To, Message string) {
	nc, err := nats.Connect("nats")
	if err != nil {
		panic(err)
	}

	fmt.Println(nc)
	ec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		panic(err)
	}
	defer ec.Close()

	sendCh := make(chan *dto.SendMessageRequest)
	ec.BindSendChan(subject, sendCh)

	i := 0
	for {
		req := &dto.SendMessageRequest{
			ChatName: subject,
			From:     From,
			To:       To,
			Message:  fmt.Sprintf("%s_%d!", Message, i),
		}
		// fmt.Printf("\n~~~~~~~~~~~~~~~~~~~~~~~\nSending request ...\nChatName:\t%s\nFrom:\t\t%s\nTo:\t\t%s\nMessage:\t%s\n~~~~~~~~~~~~~~~~~~~~~~~\n", req.ChatName, req.From, req.To, req.Message)

		// ec.Publish(subject, req)
		sendCh <- req
		time.Sleep(time.Second * 5)
		i = i + 1
	}
}

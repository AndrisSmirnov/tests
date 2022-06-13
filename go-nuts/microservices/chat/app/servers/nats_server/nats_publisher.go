package nats_server

import (
	dto "chat/api/dto"
	"fmt"
	"time"

	nats "github.com/nats-io/nats.go"
)

func Publisher(subject, From, To, Message string) {
	time.Sleep(time.Second * 5)

	i := 0
	getAnswer := true

	nc, err := nats.Connect("nats")
	if err != nil {
		panic(err)
	}

	ec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		panic(err)
	}
	defer ec.Close()

	sendCh := make(chan *dto.SendMessageRequest)
	ec.BindSendChan(subject, sendCh)

	nc.Subscribe(fmt.Sprintf("%s_response", subject), func(m *nats.Msg) {
		fmt.Printf("\n~~~~~~~~~~~~~~~~~~~~~~~\nCHAT:: Received a message: %s\n~~~~~~~~~~~~~~~~~~~~~~~\n", string(m.Data))
		nc.Close()
		getAnswer = false
	})

	for getAnswer {
		req := &dto.SendMessageRequest{
			ChatName: subject,
			From:     From,
			To:       To,
			Message:  fmt.Sprintf("%s_%d!", Message, i),
		}
		fmt.Printf("\n~~~~~~~~~~~~~~~~~~~~~~~\nUSER:: Sending a request ...\nChatName:\t%s\nFrom:\t\t%s\nTo:\t\t%s\nMessage:\t%s\n~~~~~~~~~~~~~~~~~~~~~~~\n",
			req.ChatName,
			req.From,
			req.To,
			req.Message)
		// ec.Publish(subject, req)
		sendCh <- req
		time.Sleep(time.Second * 5)
		i = i + 1
	}

	fmt.Println("CHAT:: get out")
}

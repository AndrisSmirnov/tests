package nats_server

import (
	"fmt"

	nats "github.com/nats-io/nats.go"
)

var (
	nc *nats.Conn
	ec *nats.EncodedConn
)

func Launch() {
	nc, err := nats.Connect("nats")
	if err != nil {
		panic(err)
	}

	fmt.Println(nc)
}

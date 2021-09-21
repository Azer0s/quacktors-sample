package main

import (
	"fmt"
	"github.com/Azer0s/quacktors"
	"time"
)

func main() {
	context := quacktors.RootContext()

	pingPid := quacktors.Spawn(func(ctx *quacktors.Context, message quacktors.Message) {
		fmt.Println("ping")
		<-time.After(1 * time.Second)
		ctx.Send(message.(quacktors.GenericMessage).Value.(*quacktors.Pid), quacktors.GenericMessage{Value: ctx.Self()})
	})

	pongPid := quacktors.Spawn(func(ctx *quacktors.Context, message quacktors.Message) {
		fmt.Println("pong")
		<-time.After(1 * time.Second)
		ctx.Send(message.(quacktors.GenericMessage).Value.(*quacktors.Pid), quacktors.GenericMessage{Value: ctx.Self()})
	})

	context.Send(pingPid, quacktors.GenericMessage{Value: pongPid})

	quacktors.Run()
}

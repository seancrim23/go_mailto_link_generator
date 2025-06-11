package main

import (
	"context"
	"fmt"
	"mailto_link_generator/server"
	"os"
	"os/signal"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	server, err := server.NewMailtoGeneratorServer(ctx, server.LoadConfig())
	if err != nil {
		fmt.Println("failed to create new mailto generator server: ", err)
	}

	err = server.StartMailtoGeneratorServer(ctx)
	if err != nil {
		fmt.Println("failed to start mailto generator server: ", err)
	}
}

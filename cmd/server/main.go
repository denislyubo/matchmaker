package main

import (
	"context"
	"flag"
	"log"
	"os/signal"
	"syscall"

	v1 "github.com/denislyubo/matchmaker/internal/api/v1"
	"github.com/denislyubo/matchmaker/internal/api/v1/handler"
	"github.com/denislyubo/matchmaker/internal/service"
)

func main() {
	f := flag.Uint("groupSize", 3, "Group size")
	flag.Parse()

	var groupSize uint
	if f != nil {
		groupSize = *f
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGSTOP, syscall.SIGTERM)
	defer stop()

	matchService := service.New(groupSize)
	ctl := handler.NewHandler()

	api := v1.New(matchService, ctl)

	if err := api.Start(ctx); err != nil {
		log.Fatal(err)
	}
}

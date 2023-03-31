package main

import (
	"context"
	"errors"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var (
	ErrProcessDown     = errors.New("process is down")
	ErrProcessNotFound = errors.New("process not found")
)

var logger *log.Logger

func main() {
	flags, err := parseCmdFlags()
	if err != nil {
		log.Fatal(err)
	}

	l, err := NewLogger()
	if err != nil {
		log.Fatal(err)
	}

	logger = l

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	defer cancel()

	runWithInterval(ctx, flags.Interval, processd(flags, HealthCheckLookupPid))

	wait := make(chan os.Signal, 1)
	signal.Notify(wait, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	<-wait
	logger.Println("shutting down")

	CloseLogger()
}

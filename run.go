package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

var once = new(sync.Once)

func runWithInterval(ctx context.Context, interval time.Duration, fn func() error) {
	once.Do(func() {
		fn()
	})
	ticker := time.NewTicker(interval)
	go func() {
		for {
			select {
			case <-ctx.Done():
				ticker.Stop()
				return
			case <-ticker.C:
				if err := fn(); err != nil {
					logger.Printf("error: %s", err.Error())
				}
			}
		}
	}()
}

func processd(flags Flags, healthcheck HealthCheckFunc) func() error {
	return func() error {
		err := healthcheck(flags.ProcessName)
		if err != nil {
			if errors.Is(err, ErrProcessDown) || errors.Is(err, ErrProcessNotFound) {
				return bootUp(flags, healthcheck)
			}
			return fmt.Errorf("healthcheck: %w", err)
		}

		// OK, leave it
		return nil
	}
}

func bootUp(flags Flags, h HealthCheckFunc) error {
	logger.Printf("process %s is down\n", flags.ProcessName)
	logger.Printf("running provided script: %s\n", flags.RunScript)
	if err := runScript(flags.RunScript); err != nil {
		return fmt.Errorf("run script: %w", err)
	}

	if err := h(flags.ProcessName); err != nil {
		return err
	}

	logger.Printf("process %s is up!", flags.ProcessName)
	return nil
}

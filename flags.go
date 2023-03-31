package main

import (
	"flag"
	"fmt"
	"time"
)

type Flags struct {
	Interval    time.Duration
	RunScript   string
	ProcessName string
}

func parseCmdFlags() (Flags, error) {

	interval := flag.Duration("interval", 0, "interval for healthcheck")
	script := flag.String("run-script", "", "script to start new process")
	pname := flag.String("pname", "", "process name to lookup for")
	flag.Parse()

	if *interval == time.Duration(0) {
		return Flags{}, fmt.Errorf("interval is not set")
	}

	if *script == "" {
		return Flags{}, fmt.Errorf("no script to run")
	}

	if *pname == "" {
		return Flags{}, fmt.Errorf("missing process name")
	}

	return Flags{
		Interval:    *interval,
		RunScript:   *script,
		ProcessName: *pname,
	}, nil
}

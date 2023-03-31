package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

func pidByName(pname string) (string, error) {
	cmd := exec.Command("pidof", pname)
	var buf bytes.Buffer
	cmd.Stdout = &buf
	cmd.Stderr = &buf
	if err := cmd.Run(); err != nil {
		if buf.String() == "" && strings.Contains(err.Error(), "exit status 1") {
			return "", ErrProcessNotFound
		}
		return "", fmt.Errorf("pidof unable to run: %w", err)
	}
	defer cmd.Process.Kill()
	return buf.String(), nil
}

func runScript(script string) error {
	err := exec.Command(script).Start()
	if err != nil {
		return fmt.Errorf("script unable to run: %w", err)
	}
	// Give time to boot up
	time.Sleep(time.Millisecond * 500)
	return nil
}

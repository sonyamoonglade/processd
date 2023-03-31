package main

import "fmt"

type HealthCheckFunc func(args ...any) error

func HealthCheckLookupPid(args ...any) error {
	if len(args) == 0 {
		return fmt.Errorf("invalid args. Want args[0] = {processName}")
	}
	pname, ok := args[0].(string)
	if !ok {
		return fmt.Errorf("invalid args[0] type")
	}

	_, err := pidByName(pname)
	if err != nil {
		return err
	}

	return nil
}

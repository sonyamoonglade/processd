#!/bin/bash
GOOS=linux GOARCH=amd64 go build -buildvcs=false -o processd .

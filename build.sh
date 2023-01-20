#!/bin/bash
GOOS=linux GOARCH=arm GOARM=5  go build cmd/ev3/main.go && scp main robot@ev3dev:

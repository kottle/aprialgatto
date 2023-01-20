#!/bin/bash
GOOS=linux GOARCH=arm GOARM=5  go build cmd/ev3/main.go
ssh -t robot@ev3dev 'sudo systemctl stop aprigatto.service'
scp main robot@ev3dev:
ssh -t robot@ev3dev 'sudo systemctl restart aprigatto.service'

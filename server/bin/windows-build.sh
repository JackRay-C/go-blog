#! /bin/bash

env GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build
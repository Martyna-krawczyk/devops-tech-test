#!/bin/sh

go test -short -coverprofile=coverage.out -covermode=atomic
# go tool cover -html=coverage.out
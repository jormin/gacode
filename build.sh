#!/bin/bash

# darwin
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o target/gacode_darwin_amd64_"$1" main.go

# linux
CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -o target/gacode_linux_386_"$1" main.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o target/gacode_linux_amd64_"$1" main.go
CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -o target/gacode_linux_arm_"$1" main.go

# freebsd
CGO_ENABLED=0 GOOS=freebsd GOARCH=386 go build -o target/gacode_freebsd_386_"$1" main.go
CGO_ENABLED=0 GOOS=freebsd GOARCH=amd64 go build -o target/gacode_freebsd_amd64_"$1" main.go
CGO_ENABLED=0 GOOS=freebsd GOARCH=arm go build -o target/gacode_freebsd_arm_"$1" main.go

# windows
CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o target/gacode_windows_386_"$1".exe main.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o target/gacode_windows_amd64_"$1".exe main.go
CGO_ENABLED=0 GOOS=windows GOARCH=arm go build -o target/gacode_windows_arm_"$1".exe main.go
#!/usr/bin/env bash


CGO_ENABLED=0 GOOS=windows GOARCH=386 go build  -o ./build/maker_win32/maker.exe main.go

CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build  -o ./build/maker_win64/maker.exe main.go

CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ./build/maker_mac/maker main.go

rm -f ./maker_win32.zip
rm -f ./maker_win64.zip
rm -f ./maker_mac.zip

zip -q -r maker_win32.zip ./build/maker_win32
zip -q -r maker_win64.zip ./build/maker_win64
zip -q -r maker_mac.zip ./build/maker_mac

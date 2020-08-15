#!/bin/bash

set -ev

go install -v -d github.com/therecipe/qt/cmd/...

OPWD=$PWD
unset QT_FAT
export QT_GEN_GO=true
export CGO_ENABLED=0

$(go env GOPATH)/bin/qtdeploy build darwin #same as GOOS=darwin go build -tags=minimal -ldflags="-s -w" -o ./deploy/darwin/demo
cp -r ./qml ./deploy/darwin/demo.app/Contents/MacOS/
cd $OPWD/deploy/darwin && zip -9qrXy ../demo_darwin_amd64.zip * && cd $OPWD && rm -rf $OPWD/deploy/darwin

GOOS=windows go build -tags=minimal -ldflags="-s -w -H=windowsgui" -o ./deploy/windows/demo.exe
cp -r ./qml ./deploy/windows
cd $OPWD/deploy/windows && zip -9qrXy ../demo_windows_amd64.zip * && cd $OPWD && rm -rf $OPWD/deploy/windows

GOOS=linux go build -tags=minimal -ldflags="-s -w" -o ./deploy/linux/demo
cp -r ./qml ./deploy/linux
cd $OPWD/deploy/linux && zip -9qrXy ../demo_linux_amd64.zip * && cd $OPWD && rm -rf $OPWD/deploy/linux
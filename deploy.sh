#!/bin/bash

set -ev

go install -v -tags=no_env github.com/therecipe/qt/cmd/...

OPWD=$PWD
export QT_FAT=true

$(go env GOPATH)/bin/qtdeploy -docker -tags=http_interop build windows_64_static full && docker rmi therecipe/qt:windows_64_static
cd $OPWD/full/deploy/windows && zip -9qrXy ../windows_amd64_513_full_http.zip * && cd $OPWD && rm -rf $OPWD/full/deploy/windows

$(go env GOPATH)/bin/qtdeploy -docker -tags=http_interop build linux_static full && docker rmi therecipe/qt:linux_static
cd $OPWD/full/deploy/linux && zip -9qrXy ../linux_amd64_513_full_http.zip * && cd $OPWD && rm -rf $OPWD/full/deploy/linux

cd $(go env GOPATH)/src/github.com/therecipe/qt/internal/docker/darwin && ./build_static.sh && cd $OPWD
$(go env GOPATH)/bin/qtdeploy -docker -tags=http_interop build darwin_static full && docker rmi therecipe/qt:darwin_static
cd $OPWD/full/deploy/darwin/full.app/Contents/MacOS && zip -9qrXy ../../../../darwin_amd64_513_full_http.zip * && cd $OPWD && rm -rf $OPWD/full/deploy/darwin

cd ./demo && ./deploy.sh